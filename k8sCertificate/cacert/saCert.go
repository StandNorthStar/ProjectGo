package cacert

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"k8s.io/client-go/util/keyutil"
	"log"
	"os"
	"path"
)

// 创建sa 证书
func (meta *CertMetaData) genServiceAccountKey() error {
	dir := meta.CertPath
	_, err := os.Stat(path.Join(dir, "sa.key"))
	if !os.IsNotExist(err) {
		log.Println("sa.key sa.pub already exist")
		return nil
	}
	sakey, err := NewPrivateKey(x509.RSA)
	if err != nil { return err }
	sapub := sakey.Public()

	// 写入到文件中
	err = WriteKey(dir, "sa", sakey)
	if err != nil { return err}

	return WritePublicKey(dir, "sa", sapub)
}


// 解码 sa证书
func EncodePublicKeyPEM(key crypto.PublicKey) ([]byte, error) {
	der, err := x509.MarshalPKIXPublicKey(key)
	if err != nil { return []byte{}, err}

	block := pem.Block{
		Type: PublicKeyBlockType,
		Bytes: der,
	}
	return pem.EncodeToMemory(&block), nil
}


// 保存sa证书
func WritePublicKey(pkiPath, name string, key crypto.PublicKey) error {
	if key == nil {
		return fmt.Errorf("public key cannot be nil when writing to file")
	}

	publicKeyBytes, err := EncodePublicKeyPEM(key)
	if err != nil { return err }
	publicKeyPath := pathForCert(pkiPath, name)
	if err := keyutil.WriteKey(publicKeyPath, publicKeyBytes); err != nil {
		return fmt.Errorf("unable to write public key to file %s %s", publicKeyPath, err)

	}
	return nil

}