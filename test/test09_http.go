package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
http客户端的几种方法#
1、 func (c *Client) Get(url string) (resp *Response, err error)
说明： 利用get方法请求指定的url，Get请求指定的页面信息，并返回实体主体

2、func (c *Client) Head(url string) (resp *Response, err error)
说明：利用head方法请求指定的url，Head只返回页面的首部

3、func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
说明：利用post方法请求指定的URl,如果body也是一个io.Closer,则在请求之后关闭它

4、func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
说明：利用post方法请求指定的url,利用data的key和value作为请求体.

5、func (c *Client) Do(req *Request) (resp *Response, err error)
说明：Do发送http请求并且返回一个http响应,遵守client的策略,如重定向,cookies以及auth等.当调用者读完resp.body之后应该关闭它,

如果resp.body没有关闭,则Client底层RoundTripper将无法重用存在的TCP连接去服务接下来的请求,如果resp.body非nil,则必须对其进行关闭.
通常来说,经常使用Get，Post，或者PostForm来替代Do
*/


/*
1. Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求。
2. 长连接
延申：io.reader这个是干嘛的
*/

type Data struct {
	Msg string 		`json:"msg"`
	Code int 		`json:"code"`
	Request string 	`json:"request"`
}
type HarborData struct {
	Project_name string 	`json:"project_name"`
	Public int 				`json:"public"`
}

func main() {
	/*
	-XGET
	1. Get Body
	2. 指定账号和密码
		req1.SetBasicAuth("admin", "Harbor@0303!")
	   	curl -XGET -u 'admin:Harbor@0303!'  -H "Content-Type: application/json" "http://106.37.75.67/api/projects"
	3. 带参数获取相关数据
	 */

	// 1. Get获取页面信息
	fmt.Println("1. Get获取页面信息--------------------------------------------------")
	url := "https://api.douban.com/v2/movie/search"
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	log.Printf("GET1-1: %T \n", req)
	log.Println("GET1-1: ",req.Body)
	// 输出get到的body结果
	log.Println(string(body))

	// 把json结果解析为go结构体
	var rd Data
	r := json.Unmarshal(body, &rd)
	if r != nil {
		log.Println("GET1-1: json解析失败",r)
	}
	log.Println("GET1-1: Trasfer Go Type: ",rd)
	fmt.Println("--------------------------------------------------")

	// 2. 通过 用户名和密码验证get请求
	/*
	方式一 使用http.NewRequest
	先生成http.client -> 再生成 http.request -> 之后提交请求：client.Do(request) -> 处理返回结果，每一步的过程都可以设置一些具体的参数
	 */
	fmt.Println("2. 通过 用户名和密码验证get请求--------------------------------------------------")
	url1 := "http://xxx.xxx.xxx.xxx/api/projects?project_name=xxx"
	//生成client 参数为默认
	client1 := &http.Client{}
	//提交请求
	req1, _err := http.NewRequest("GET", url1, nil)
	if _err != nil {
		log.Println("GET1-2: Error1 ",_err)
	}
	//设置request的header
	req1.SetBasicAuth("xxx", "xxx@xxx!")
	req1.Header.Set("Content-Type", "Application/json")

	req2, _err2 := client1.Do(req1)
	if _err2 != nil {
		log.Println("GET1-2: ERROR2 ", _err2)
	}
	defer req2.Body.Close()
	log.Println("GET1-2: ", req2.StatusCode)
	log.Println("GET1-2: ", req2.Body)
	body2, _ := ioutil.ReadAll(req2.Body)
	log.Println("GET1-2: ", string(body2))
	fmt.Println("--------------------------------------------------")


	// 3. GET带参数请求   (暂时未实现)


	// 4. POST创建请求
	/*
	post
	*/
	fmt.Println("4. POST创建请求--------------------------------------------------")
	url2 := "http://xxx.xxx.xxx.xxx/api/projects"
	hd := HarborData{
		"test01",
		0,
	}
	hd_data, err := json.Marshal(hd)
	if err != nil {
		fmt.Println("POST2-1: ERROR1 ", err)
	}
	fmt.Println("POST2-1: ", string(hd_data))
	fmt.Printf("POST2-1: TYPE %T \n",hd_data)

	client2 := &http.Client{}
	req3, err := http.NewRequest("POST", url2, strings.NewReader(string(hd_data)))
	if err != nil {
		fmt.Println("POST2-1: ERROR2 ", err)
	}
	req3.SetBasicAuth("xxx", "xxx@xxx!")
	req4, err := client2.Do(req3)
	if err != nil {
		fmt.Println("POST2-1: ERROR3 ", err)
	}
	defer req4.Body.Close()
	// body3, _ := ioutil.ReadAll(req4.Body)
	fmt.Println("POST2-1: REQ4 RESUTL MEM ADDR ", req4.Body)
	fmt.Println("--------------------------------------------------")

	/*
	自定义客户端
	http transport  RoundTrip
	*/



	/*
	长连接
	 */
}

