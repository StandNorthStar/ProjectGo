package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)


type Alerts struct {
	Annotations map[string]interface{} `json:annotations`
	StartsAt    time.Time   `json:startsAt`
	EndsAt      time.Time   `json:endsAt`
	Status      string      `json:status`
	Labels      map[string]interface{}     `json:labels`
}

type AlertsData struct {
	Status string   `json:status`
	Alerts []Alerts `json:alerts`
}

type Params struct {
	Url   string
	Path  string
	Proxy string
}

const templateText = `
{{- if eq .Status "firing" }}
    <font color=red size=20>告警</font>\n
    名称： {{ .Labels.alertname }}\n
    描述： {{ .Annotations.description }}\n
    地址： {{ .Labels.instance }}\n
    告警值： {{ .Annotations.value }}\n
    开始时间： {{ (.StartsAt.Add 28800e9).Format "2006-01-02 15:04:05" }}
{{- else if eq .Status "resolved" }}
    <font color=#228b22 size=20>恢复</font>\n
    名称： {{ .Labels.alertname }}\n
    描述： {{ .Annotations.description }}\n
    地址： {{ .Labels.instance }}\n
    告警值： {{ .Annotations.value }}\n
    开始时间： {{ (.StartsAt.Add 28800e9).Format "2006-01-02 15:04:05" }}\n
    结束时间： {{ (.EndsAt.Add 28800e9).Format "2006-01-02 15:04:05" }}
{{- end }}`

// 发送消息类型
func MsgMarkdown(msg string) string {

	return fmt.Sprintf(`{
		"msgtype": "markdown",
		"markdown": {
            "content": "%s"
		}
	}}`, msg)
}

func MsgTemplate(templatePath string, m Alerts) (string, error) {

	if len(templatePath) == 0 {

		temp := template.Must(template.New("anyname").Parse(templateText))
		var content bytes.Buffer
		contentErr := temp.Execute(&content, m)
		return content.String(), contentErr
	}

	// 判断模板文件是否存在
	_, err := os.Stat(templatePath)
	if os.IsNotExist(err) {
		log.Fatalf("File %s Not Exsit", templatePath)
	}

	temp, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	var content bytes.Buffer
	contentErr := temp.Execute(&content, m)
	fmt.Println("result: ", content.String())
	return content.String(), contentErr

}

func WeChatSend(url string, data io.Reader, https_proxy string) (*http.Response, error) {

	// 如设置代理，添加代理变量
	if len(https_proxy) != 0 {
		os.Setenv("https_proxy", https_proxy)
	}

	client := &http.Client{
		// 设置http 代理
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	prereq, _err := http.NewRequest("POST", url, data)
	if _err != nil {
		log.Println(_err)
	}
	prereq.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, _err := client.Do(prereq)
	return response, _err


}

func (p Params) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(body))

	var alertdata AlertsData
	if err := json.NewDecoder(r.Body).Decode(&alertdata); err != nil {
		log.Println("request data error: ",err)
		//log.Fatal(err)
	}
	log.Println("--- json parse data start ---")
	log.Println(alertdata)
	log.Println("--- json parse data end ---")

	alert_alerts := alertdata.Alerts
	for _, v := range alert_alerts {

		msg, err := MsgTemplate(p.Path, v)
		if err != nil {
			//log.Fatal(err)
			log.Println("Parsh Template Error:", err)
		}
		msgmarkdown := MsgMarkdown(msg)

		DataByte := []byte(msgmarkdown)
		DataReader := bytes.NewReader(DataByte)
		// 异常处理 recover
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		req, reqerr := WeChatSend(p.Url, DataReader, p.Proxy)
		if reqerr != nil {
			log.Println(reqerr)
		}
		req.Body.Close()

	}
	log.Println(alertdata.Status, alert_alerts)
}

func main() {

	webhook := pflag.StringP("webhook", "w", "", "enterprise wechat webhook url.")
	templatePath := pflag.StringP("template", "t", "", "Alert Content Template, if not setting use defaults.")
	port := pflag.IntP("port", "p", 5001, "Prometheus Webhook Listen Port. ")
	proxy := pflag.String("proxy", "", "Alert Proxy.")
	pflag.Parse()

	if len(*webhook) <= 0 {
		panic("Please Setting Webhook URL")
	}

	if templatePath != nil {
		fmt.Println(*templatePath)
	}

	parms := Params{
		Url:   *webhook,
		Path:  *templatePath,
		Proxy: *proxy,
	}

	mux := http.NewServeMux()
	mux.Handle("/alert", parms)

	// 启动
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", strconv.Itoa(*port)),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
