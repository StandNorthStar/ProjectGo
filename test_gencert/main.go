package main

import (
	"fmt"
	"github.com/wonderivan/logger"
	. "test_gencert/cert"

)

func main() {
	BasePath := "/mnt/d/GoProject/test_gencert/gencert"
	EtcdBasePath := "/mnt/d/GoProject/test_gencert/gencert/etcd"

	fmt.Println("开始生成证书---")

	certMeta, err := NewSealosCertMetaData(BasePath,
		EtcdBasePath,
		[]string{"xxx.com", "10.56.xxx.xxx", "kubernetes.default.svc.sealyun"},
		"172.64.0.0/10",
		"xxx",
		"xxx.xxx.xxx.xxx",
		"cluster.local")
	if err != nil {
		//logger.Info("creating kubeconfig file for %s", kubeConfigFileName)
		logger.Error(err)
	}

	if err := certMeta.GenerateAll() ; (err != nil ) {
		logger.Info("Ok ")
	}

}
