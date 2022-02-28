package main

import (
	"fmt"
	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	yaml "gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Endpoint 		string
	AccessKeyID		string
	AccessKeySecret	string
	ProxyHost		string
}


func main() {
	// 解析配置
	type ObjProvider string
	const (
		FILESYSTEM ObjProvider = "FILESYSTEM"
		GCS        ObjProvider = "GCS"
		S3         ObjProvider = "S3"
		AZURE      ObjProvider = "AZURE"
		SWIFT      ObjProvider = "SWIFT"
		COS        ObjProvider = "COS"
		ALIYUNOSS  ObjProvider = "ALIYUNOSS"
	)
	type BucketConfig struct {
		Type   ObjProvider `yaml:"type"`
		Config interface{} `yaml:"config"`
	}

	confContentYaml, _ := os.ReadFile("/home/heliyun/aa/oss.yaml")

	bucketConf := &BucketConfig{}
	if err := yaml.UnmarshalStrict(confContentYaml, bucketConf); err != nil {
		fmt.Println(err)
	}
	config_v1, err := yaml.Marshal(bucketConf.Config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config_v1)

	fmt.Println("---\n")
	fmt.Println("---\n")
	fmt.Println("---\n")
	fmt.Println("---\n")


	var config Config
	config = Config{
		Endpoint:        "oss-x-x-x-x-x.aliyuncs.com",
		AccessKeyID:     "x",
		AccessKeySecret: "x",
		ProxyHost:       "http://x.x.x.x:8000",
	}

	client, err := alioss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret, alioss.Proxy(config.ProxyHost))
	//client, err := alioss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	bucketname := "oss-bucketname"
	//bk, err := client.Bucket(config.Bucket)
	bk, err := client.Bucket(bucketname)
	if err != nil {
		fmt.Println(err)
	}

	marker := ""
	for {
		lsRes, err := bk.ListObjects(alioss.Marker(marker))
		if err != nil {
			fmt.Println(err)
		}
		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			fmt.Println("Bucket: ", object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}

}




