package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"
	"time"
)

type Annotations struct {
	Describe string `json:describe`
	Summary  string `json:summary`
	Value    string `json:value`
}
type Labels struct {
	Alertname string `json:alertname`
	Instance  string `json:instance`
	Job       string `json:job`
	Severity  string `json:severity`
}

type Alerts struct {
	Annotations Annotations `json:annotations`
	StartsAt   time.Time  	`json:startsAt`
	EndsAt     time.Time  	`json:endsAt`
	Status     string     	`json:status`
	Labels     Labels       `json:labels`
}

func main() {

	a := `{
		"status": "firing",
		"labels": {
			"alertname": "go version",
			"instance": "localhost:9090",
			"job": "prometheus",
			"severity": "warning",
			"version": "go1.16.7"
		},
		"annotations": {
			"describe": "HOST: localhost:9090",
			"summary": "go is version 1.16",
			"value": "1"
		},
		"startsAt": "2021-12-20T10:03:53.909Z",
		"endsAt": "0001-01-01T00:00:00Z",
		"generatorURL": "http://LAPTOP-9M889QU5:9090/graph?g0.expr=go_info%7Bjob%3D%22prometheus%22%7D+%3D%3D+1\u0026g0.tab=1",
		"fingerprint": "8befb7bbd757915b"
	}`
	var alerts Alerts
	err := json.Unmarshal([]byte(a), &alerts)
	if err != nil {
		log.Fatal(err)
	}

	tempath := "./default.tmpl"
	temp, err := template.ParseFiles(tempath)
	if err != nil {
		log.Fatal(err)
	}

	//输出1： 把模板内容输出到终端
	//err1 := temp.Execute(os.Stdout, alerts)

	//输出2： 把模板内容输出成字符串
	var buf bytes.Buffer
	err1 := temp.Execute(&buf, alerts)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("---")
	fmt.Println(buf.String())
}

