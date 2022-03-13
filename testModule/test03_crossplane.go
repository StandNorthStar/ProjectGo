package main

import (
	"bytes"
	"fmt"
	"github.com/aluttik/go-crossplane"
	"log"
)

func main() {
	// 查看配置
	filename := "/usr/local/openresty/nginx/conf/nginx.conf"
	//filename := "/usr/local/openresty/nginx/conf/vhost/test.conf"
	searchconfig(filename)

	/*
		// 新增配置
		filename := "/home/heliyun/aa.conf"
		content := `
		server {
			listen 443 ssl;
			server_name test03-json.class.com;
			ssl_certificate key/dev1/dev1_673.pem;
			ssl_certificate_key key/dev1/dev1_673.key;
			access_log logs/jsaq.weixin-access.log reverseRealIpFormat1;
			error_log logs/jsaq.weixin-error.log;
			keepalive_timeout 60;
			error_page 404 500 502 503 504 /50x.html;
			client_max_body_size 60m;
			location / {
			access_log logs/test-access-loc.log reverseRealIpFormat1;
			proxy_pass https://localhost:8080;
		}
		}
		server {
			listen 80;
			server_name test03-json.class.com;
			access_log logs/jsaq.weixin-access.log reverseRealIpFormat1;
			error_log logs/jsaq.weixin-error.log;
			keepalive_timeout 60;
			error_page 404 500 502 503 504 /50x.html;
			client_max_body_size 60m;
			location / {
			access_log logs/test-access-loc.log reverseRealIpFormat1;
			proxy_pass https://localhost:8080;
		}
		}
		`
		addconfig(filename, content)
	*/

}

/*
导出nginx配置
*/
func dump(conf crossplane.Config) {

	var buf bytes.Buffer
	if err := crossplane.Build(&buf, conf, &crossplane.BuildOptions{}); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

/*
查看配置
*/
func searchconfig(filename string) {
	payload, err := crossplane.Parse(filename, &crossplane.ParseOptions{})
	if err != nil {
		log.Println(err)
	}

	for _, i := range payload.Config {
		fmt.Println()
	}
	//b, err := json.Marshal(payload)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(b))
	//fmt.Println(payload.Status)
	//
	//for _, v := range payload.Config {
	//
	//	fmt.Println("--------------------")
	//	fmt.Println(v.File)
	//	//fmt.Println(v.Parsed)
	//	//fmt.Println("\n")
	//	dump(v)
	//}
}

func addconfig(filename, content string) {
	//crossplane.Build()
}
