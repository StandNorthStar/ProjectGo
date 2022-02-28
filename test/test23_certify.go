package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

/*
golang的x509标准库下有个Certificate结构,这个结构就是证书解析后对应的实体.新证书需要先生成秘钥对,然后使用根证书的私钥进行签名.证书和私钥以及公钥这里使用的是pem编码方式.

公钥，私钥和数字签名
公钥和私钥是成对的，它们互相解密。
公钥加密: 需要私钥解密。
私钥加密：就是私钥数字签名，使用公钥验证。

只有pfx格式的数字证书是包含有私钥的，cer格式的数字证书里面只有公钥没有私钥
证书参考：https://www.cnblogs.com/xq1314/archive/2017/12/05/7987216.html
 */
func ParsePem() {
	fmt.Println("----- 读取根证书的证书和私钥")
	filepath, err := os.Getwd()
	if err != nil {
		fmt.Println("getwd error:", err)
		return
	}
	// 解析公钥
	rootCa := filepath + "/ca.crt" // ca证书公钥
	fmt.Println(rootCa)
	caFile, err := ioutil.ReadFile(rootCa)
	if err != nil {
		fmt.Println("readfile error: ",err)
		return
	}
	caBlock, _ := pem.Decode(caFile)

	cert, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		fmt.Println("x509 Parse error:", err)
		return
	}
	fmt.Println(cert)

	// 解析私钥
	rootKey := filepath + "/ca.key"   // ca证书私钥
	keyFile, err := ioutil.ReadFile(rootKey)
	if err != nil {
		return
	}
	keyBlock, _ := pem.Decode(keyFile)
	praKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return
	}
	fmt.Println(praKey)
}

func GenCa() {

	serialNumber := new(big.Int).SetInt64(0)
	duration365d := time.Hour * 24 * 365
	year := time.Duration(100)
	template := &x509.Certificate{
		SerialNumber:   serialNumber, 	// SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		Subject:        pkix.Name{   	//Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略。
							Organization:       []string{"xxx.xxx."},
							CommonName:         "k8s app",
						},
		NotBefore:      time.Now(),			                   //证书有效期开始时间
		NotAfter:       time.Now().Add(duration365d * year),   //证书有效期结束时间
		KeyUsage:       x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, //KeyUsage 与 ExtKeyUsage 用来表明该证书是用来做服务器认证的
		BasicConstraintsValid: true,				//基本的有效性约束
		IsCA:                  true,				 //是否是根证书
		//ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, // 密钥扩展用途的序列  (证书用途(客户端认证，数据加密))
		//IPAddresses:    []net.IP{net.ParseIP("127.0.0.1")},
	}

	//生成公钥私钥对
	/*
	rsa.GenerateKey 使用随机源随机（例如，crypto/rand.Reader）生成给定位大小的 RSA 密钥对。
	--- 创建密钥对 ---
	 */
	key, _ := rsa.GenerateKey(rand.Reader, 2048) //生成一对具有指定字位数的RSA密钥

	/*
	x509.CreateCertificate 根据模板创建一个新的证书; 返回的片是 DER 编码中的证书。
	支持通过 crypto.Signer 实现的所有密钥类型（包括 *rsa.PublicKey 和 *ecdsa.PublicKey。）
	参数说明：1.Reader;2.证书模板；3.证书模板；4.公钥；5.私钥
	注意：如果 x509.CreateCertificate 的第二个和第三个参数一样，那么这个是自签名证书。（其实第三个参数是父模板）
	--- 这一步其实是给创建的公钥和私钥签名 ---
	*/
	certDERBytes, err := x509.CreateCertificate(rand.Reader, template, template, key.Public(), key)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("-----------------")
	fmt.Println(certDERBytes)
	re001, err := x509.ParseCertificate(certDERBytes)
	fmt.Println("----------")
	fmt.Println(re001)
	fmt.Println("-----------------")

	// 输出公钥
	// pem.Block encoding实现了PEM数据编码，可以对密钥做输出和输入
	block := pem.Block{
		//Type:  "CertificateBlockType",
		Type:  "CERTIFICAET",
		Bytes: certDERBytes,
	}
	fmt.Println("--output public key --")
	//fmt.Println(pem.EncodeToMemory(&block))
	fmt.Println(string(pem.EncodeToMemory(&block)))

	// 输出私钥
	block_private := pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	fmt.Println("--output private key --")
	//fmt.Println(pem.EncodeToMemory(&block_private))
	fmt.Println(string(pem.EncodeToMemory(&block_private)))


	fmt.Println("------------")
	re, err := x509.ParseCertificate(certDERBytes)
	fmt.Println(re.PublicKey)
	//fmt.Println(string(re.Raw))

	// 把结果输出到文件
	//pk, _ := rsa.GenerateKey(rand.Reader, 2048) //生成一对具有指定字位数的RSA密钥
	//CreateCertificate基于模板创建一个新的证书
	//第二个第三个参数相同，则证书是自签名的
	//返回的切片是DER编码的证书
	//derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) //DER 格式
	//certOut, _ := os.Create("cert.pem")
	//pem.Encode(certOut,&pem.Block{Type:"CERTIFICAET", Bytes: derBytes})
	//certOut.Close()
	//keyOut, _ := os.Create("key.pem")
	//pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	//keyOut.Close()
}

func main() {
	GenCa()
}


