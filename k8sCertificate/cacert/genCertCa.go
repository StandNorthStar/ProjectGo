package cacert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"

	//"net"
	"time"
)

const (
	// PrivateKeyBlockType is a possible value for pem.Block.Type. PKCS#8格式
	PrivateKeyBlockType = "PRIVATE KEY"
	// PublicKeyBlockType is a possible value for pem.Block.Type.  PKCS#8格式 (sa证书)
	PublicKeyBlockType = "PUBLIC KEY"
	// CertificateBlockType is a possible value for pem.Block.Type. CER
	CertificateBlockType = "CERTIFICATE"
	// RSAPrivateKeyBlockType is a possible value for pem.Block.Type. CER
	RSAPrivateKeyBlockType = "RSA PRIVATE KEY"
	rsaKeySize             = 2048
	duration365d           = time.Hour * 24 * 365
)


// AltNames contains the domain names and IP addresses that will be added
// to the API Server's x509 certificate SubAltNames field. The values will
// be passed directly to the x509.Certificate object.
// x509 认证DNS域信息
type AltNames struct {
	DNSNames map[string]string
	IPs      map[string]net.IP
}
type Config struct {
	Path         string // Writeto Dir
	DefaultPath  string // Kubernetes default Dir
	BaseName     string // Writeto file name
	CAName       string // root ca map key
	CommonName   string
	Organization []string
	Year         time.Duration
	AltNames     AltNames
	Usages       []x509.ExtKeyUsage
}



func NewCaCertAndKey(cfg Config) (*x509.Certificate, crypto.Signer, error) {
	// 1. 判断本地是否存在密钥
	_, err := os.Stat(pathForKey(cfg.Path, cfg.BaseName))
	if !os.IsNotExist(err) {
		fmt.Println("ca文件存在返回")
		return nil,nil,fmt.Errorf("%s 文件已经存在 : %s",cfg.BaseName, err)
	}
	// 2. 创建随机的密钥对
	key, err := NewPrivateKey(x509.UnknownPublicKeyAlgorithm)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create private key while generating CA certificate %s", err)
	}

	// 3. 根据设置的信息给此密钥对签名
	cert, err := NewSelfSignedCACert(key, cfg.CommonName, cfg.Organization, cfg.Year)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create ca cert %s", err)
	}
	return cert, key, nil
}

//生成随机密钥对
func NewPrivateKey(keyType x509.PublicKeyAlgorithm) (crypto.Signer, error) {
	if keyType == x509.ECDSA {
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	}
	return rsa.GenerateKey(rand.Reader, rsaKeySize)
}

//为密钥对签名
func NewSelfSignedCACert(key crypto.Signer, commonName string, organization []string, year time.Duration) (*x509.Certificate, error) {
	now := time.Now()
	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: organization,
		},
		NotBefore:             now.UTC(),
		NotAfter:              now.Add(duration365d * year).UTC(),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	//  x509.CreateCertificate(rand.Reader, 被签名证书信息，父证书信息，被签名公钥，被签名私钥)
	certDERBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		return nil, err
	}
	return x509.ParseCertificate(certDERBytes)
}


func pathForCert(pkiPath, name string) string {
	return filepath.Join(pkiPath, fmt.Sprintf("%s.crt", name))
}
func pathForKey(pkiPath, name string) string {
	return filepath.Join(pkiPath, fmt.Sprintf("%s.key", name))
}



