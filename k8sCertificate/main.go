//package k8sCertificate
package main

import (
	"k8sCertificate/cacert"
	"log"
	"net"
)
/*
问题：
1. 生成证书时添加的DNS和IP地址是循环还是一次性生成的？
2.
 */
func main() {
	apidns := map[string]string{}
	apidns["t1.xxx.com"] = "t1.xxx.com"
	apidns["t2.xxx.com"] = "t2.xxx.com"

	apiips := map[string]net.IP{}
	apiips["192.168.16.105"] = net.IPv4(192,168,16,105)
	apiips["192.168.16.106"] = net.IPv4(192,168,16,106)
	apiips["192.168.16.107"] = net.IPv4(192,168,16,107)


	altname := cacert.AltNames{
		DNSNames: apidns,
		IPs: apiips,
	}
	parms := cacert.CertMetaData{
		APIServer: altname,
		NodeName: "k8s-test-01",
		NodeIP: "192.168.16.105",
		DNSDomain: "xxx.com",

		CertPath: "/home/heliyun/.k8s01",
		CertEtcdPath: "/home/heliyun/.k8s01/etcd",
	}

	/*
	InitCertMetaData(CertPath, CertEtcdPath string, apiServiceIPDomains []string, SvcCIDR, nodeName, nodeIP, DNSDomain string)
	 */

	err := parms.GenCertMain()
	if err != nil {
		log.Println("Gen Cert Failed Main!")
	}

}


