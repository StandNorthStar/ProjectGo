package main

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	k8scmdapi "k8s.io/client-go/tools/clientcmd/api"
)


func main() {

	cluster := k8scmdapi.Cluster{
		Server:    			"https://localhost-test:6443",  // 集群地址
		CertificateAuthorityData:  []byte{'1','2','N'},     // pem格式的ca根证书，用于https

	}

	authInfo := k8scmdapi.AuthInfo{
		ClientCertificateData:	[]byte{'c','e','r','d'},  // pem格式的用户证书
		//ClientKey:			"test",
		ClientKeyData:		[]byte{'k','e','y','d'},	  // pem格式的用户私钥
		//TokenFile:			"admin.conf",

	}

	context := k8scmdapi.Context{
		//LocationOfOrigin: "context",
		Cluster: "kubernets",
		AuthInfo: "kubernetes-admin",
		//Namespace: "defaults",
		//Extensions: map[string]runtime.Object{},
	}

	k8sconfig := k8scmdapi.Config{
		Kind: 			"Kind",
		APIVersion:  	"v1",
		Preferences: 	k8scmdapi.Preferences{},   // 为空
		Clusters: map[string]*k8scmdapi.Cluster{
			"kubernetes": &cluster,
		},
		AuthInfos: map[string]*k8scmdapi.AuthInfo{  // 用户信息, 所以你直接改kubeconfig里的user是没用的，因为k8s只认证书里的名字
			"kubernetes-admin":	&authInfo,
		},
		Contexts: map[string]*k8scmdapi.Context{
			"kubernetes-admin@kubernetes": &context,
		},
		CurrentContext: "kubernetes-admin@kubernetes",   // 当前上下文, kubeconfig可以很好支持多用户和多集群
		//Extensions: map[string]runtime.Object{},
	}

	err := WriteToDisk("/home/xxx/.k8s01/k8s-admin-02.conf", &k8sconfig)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}


// WriteToDisk writes a KubeConfig object down to disk with mode 0600
func WriteToDisk(filename string, kubeconfig *k8scmdapi.Config) error {
	err := clientcmd.WriteToFile(*kubeconfig, filename)
	if err != nil {
		return err
	}
	return nil
}

func LoadCert(filename string) {
	filepath := "/home/xxx/.k8s01"
	fmt.Println(filepath)

}

func createCertFile() {

}


