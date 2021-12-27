package cacert

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"log"
	"net"
)

var (
	KubernetesDir			= "/home/xxx/.k8s"
	KubeDefaultCertPath 	= "/home/xxx/.k8s/pki"
	kubeDefaultCertEtcdPath = "/home/xxx/.k8s/pki/etcd"
)

func CaList(CertPath, CertEtcdPath string) []Config {
	return []Config{
		{
			Path:			CertPath,
			DefaultPath:	KubeDefaultCertPath,
			BaseName:		"ca",
			CommonName:		"kubernetes",
			Organization:	nil,
			Year:			100,
			AltNames:		AltNames{},
			Usages:			nil,
		},
		{
			Path: 			CertPath,
			DefaultPath: 	KubeDefaultCertPath,
			BaseName:  		"front-proxy-ca",
			CommonName:    	"front-proxy-ca",
			Organization:	nil,
			Year:			100,
			AltNames:		AltNames{},
			Usages:			nil,
		},
		{
			Path:  			CertEtcdPath,
			DefaultPath:  	kubeDefaultCertEtcdPath,
			BaseName:  		"etcd",
			CommonName:  	"etcd-ca",
			Organization:  	nil,
			Year:  			100,
			AltNames:  		AltNames{},
			Usages:  		nil,
		},
	}
}

func CertList(CertPath, CertEtcdPath string) []Config {
	return []Config{
		{
			Path: 			CertPath,
			DefaultPath:  	KubeDefaultCertPath,
			BaseName:		"apiserver",
			CAName:  		"kubernetes",
			CommonName:   	"kube-apiserver",
			Organization:  	nil,
			Year:  			100,
			AltNames: 		AltNames{
				DNSNames: map[string]string{
					"localhost":				"localhost",
					"kubernetes":				"kubernetes",
					"kubernetes.default":		"kubernetes.default",
					"kubernetes.default.svc":	"kubernetes.default.svc",
				},
				IPs: map[string]net.IP{
					"127.0.0.1":	net.IPv4(127,0, 0, 1),
				},
			},
			Usages:			[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		},
		{
			Path:         CertPath,
			DefaultPath:  KubeDefaultCertPath,
			BaseName:     "apiserver-kubelet-client",
			CAName:       "kubernetes",
			CommonName:   "kube-apiserver-kubelet-client",
			Organization: []string{"system:masters"},
			Year:         100,
			AltNames:     AltNames{},
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		},
		{
			Path:         CertPath,
			DefaultPath:  KubeDefaultCertPath,
			BaseName:     "front-proxy-client",
			CAName:       "front-proxy-ca",
			CommonName:   "front-proxy-client",
			Organization: nil,
			Year:         100,
			AltNames:     AltNames{},
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		},
		{
			Path:         CertPath,
			DefaultPath:  KubeDefaultCertPath,
			BaseName:     "apiserver-etcd-client",
			CAName:       "etcd-ca",
			CommonName:   "kube-apiserver-etcd-client",
			Organization: []string{"system:masters"},
			Year:         100,
			AltNames:     AltNames{},
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		},
		{
			Path:         CertEtcdPath,
			DefaultPath:  kubeDefaultCertEtcdPath,
			BaseName:     "server",
			CAName:       "etcd-ca",
			CommonName:   "etcd", // kubeadm using node name as common name cc.CommonName = mc.NodeRegistration.Name
			Organization: nil,
			Year:         100,
			AltNames:     AltNames{}, // need set altNames
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		},
		{
			Path:         CertEtcdPath,
			DefaultPath:  kubeDefaultCertEtcdPath,
			BaseName:     "peer",
			CAName:       "etcd-ca",
			CommonName:   "etcd-peer", // change this in filter
			Organization: nil,
			Year:         100,
			AltNames:     AltNames{}, // change this in filter
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		},
		{
			Path:         CertEtcdPath,
			DefaultPath:  kubeDefaultCertEtcdPath,
			BaseName:     "healthcheck-client",
			CAName:       "etcd-ca",
			CommonName:   "kube-etcd-healthcheck-client",
			Organization: []string{"system:masters"},
			Year:         100,
			AltNames:     AltNames{},
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		},
	}
}

// 接收的参数
type CertMetaData struct {
	APIServer AltNames	// apiserver 需要结束的参数
	NodeName  string	// etcd证书需要的参数
	NodeIP    string	// etcd证书需要参数
	DNSDomain string

	//证书生成的位置
	CertPath     string
	CertEtcdPath string
}

const (
	APIserverCert = iota
	EtcdServerCert
	EtcdPeerCert
)

/*
apiserver : 所有Master节点IP + VIP + ServiceIP第一个IP + kubernetes域名 + 127.0.0.1
Etcd: 所有etcd节点IP+ 127.0.0.1
作用：初始化 CertMetaData 数据
 */
func InitCertMetaData(CertPath, CertEtcdPath string, apiServiceIPDomains []string, SvcCIDR, nodeName, nodeIP, DNSDomain string) (*CertMetaData, error) {

	data := &CertMetaData{}
	data.CertPath = CertPath
	data.CertEtcdPath = CertEtcdPath
	data.DNSDomain = DNSDomain

	data.APIServer.IPs = make(map[string]net.IP)
	data.APIServer.DNSNames = make(map[string]string)

	svcFirstIP, _, err := net.ParseCIDR(SvcCIDR)
	if err != nil { return nil, err }
	svcFirstIP[len(svcFirstIP)-1]++ // 获取svc第一个IP
	data.APIServer.IPs[svcFirstIP.String()] = svcFirstIP

	for _, altname := range apiServiceIPDomains {
		// apiServiceIPDomains既有IP地址，也有域名
		ip := net.ParseIP(altname)
		if ip != nil {  // 当解析域名时，ip为nil
			data.APIServer.IPs[ip.String()] = ip
			continue    // 当获取的是IP地址时，已经赋值跳出循环，下面域名的赋值直接跳过。
		}
		data.APIServer.DNSNames[altname] = altname
	}

	ip := net.ParseIP(nodeIP)
	if ip != nil {
		data.APIServer.IPs[ip.String()] = ip
	}

	data.NodeIP = nodeIP
	data.NodeName = nodeName
	return data, nil
}



func (meta CertMetaData) modifyApiserverAltNames(certList *[]Config) {

	// 传递的认证IP和DNS替换掉原有的数据
	/*
	待梳理
	 */
	for _, dns := range meta.APIServer.DNSNames {
		//(*certList)[APIServer]
		(*certList)[APIserverCert].AltNames.DNSNames[dns] = dns
	}

	svcDNS := fmt.Sprintf("kubernetes.default.svc.%s", meta.DNSDomain)
	(*certList)[APIserverCert].AltNames.DNSNames[svcDNS] = svcDNS
	(*certList)[APIserverCert].AltNames.DNSNames[meta.NodeName] = meta.NodeName

	for _, ip := range meta.APIServer.IPs {
		(*certList)[APIserverCert].AltNames.IPs[ip.String()] = ip
	}

	log.Printf("ApiServer AltNames: %v", (*certList)[APIserverCert].AltNames)
}

func (meta CertMetaData) modifyEtcdAltNames(certList *[]Config) {
	altnames := AltNames{
		DNSNames: map[string]string{
			"localhost": "localhost",
			meta.NodeName: meta.NodeName,
		},
		IPs: map[string]net.IP{
			net.IPv4(127,0,0,1).String(): 	net.IPv4(127,0,0,1),
			net.ParseIP(meta.NodeIP).To4().String():  	net.ParseIP(meta.NodeIP).To4(),
			net.IPv6loopback.String(): 					net.IPv6loopback,

		},
	}
	(*certList)[EtcdServerCert].CommonName = meta.NodeName
	(*certList)[EtcdServerCert].AltNames = altnames
	(*certList)[EtcdPeerCert].CommonName = meta.NodeName
	(*certList)[EtcdPeerCert].AltNames = altnames

	log.Printf("Etcd altnames: %v, commonName: %s", (*certList)[EtcdPeerCert].AltNames, (*certList)[EtcdPeerCert].CommonName)
}

func (meta *CertMetaData) GenCertMain() error {
	cas := CaList(meta.CertPath, meta.CertEtcdPath)
	certs := CertList(meta.CertPath, meta.CertEtcdPath)
	meta.modifyApiserverAltNames(&certs)
	meta.modifyEtcdAltNames(&certs)
	meta.genServiceAccountKey()

	/*
	定义初始化两个map类型的变量: CACerts、CAKeys
	意义：把CA证书的公钥和私钥存储到这两个变量中，以便为其他服务证书签名。
	CaList.Config.CommonName == CertList.Config.CAName
	 */
	CACerts := map[string]*x509.Certificate{}
	CAKeys := map[string]crypto.Signer{}

	for _, ca := range cas {
		caCert, cakey, err := NewCaCertAndKey(ca)
		if err != nil { return err }

		// 为CACerts、CAKeys赋值
		CACerts[ca.CommonName] = caCert
		CAKeys[ca.CommonName] = cakey

		err = WriteCertAndKey(ca.Path, ca.BaseName, caCert, cakey)
		if err != nil { return err }
	}

	for _, cert := range certs {
		//map 类型元素查找；判断是否成功找到特定的键，不需要检查取到的值是否为nil，只需查看第二个返回值。
		caCert, ok := CACerts[cert.CAName]
		if !ok {
			return fmt.Errorf("root ca cert not found %s", cert.CAName)
		}
		caKey, ok := CAKeys[cert.CAName]
		if !ok {
			return fmt.Errorf("root ca key not found %s", cert.CAName)
		}

		Cert, Key, err := NewCaCertAndKeyFromRoot(cert, caCert, caKey)
		if err != nil {
			return err
		}

		err = WriteCertAndKey(cert.Path, cert.BaseName, Cert, Key)
		if err != nil {
			return err
		}
	}

	return nil
}


