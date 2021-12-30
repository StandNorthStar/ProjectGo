/*
 xxl-job version : 1.9.0
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	httpurl "net/url"
	"os"
	"strconv"
	"time"
)

type appInfo struct {
	AppId   string
	AppName string
	AppDesc string
}

type jobInfo struct {
	JobId     int    `json:"id"`
	JobDesc   string `json:"jobDesc"`
	JobStatus string `json:"jobStatus"`
}

//type TaskLog struct {
//	Id			int 		`json:"id"`
//	JobGroup 	int 		`json:"jobGroup"`
//	JobId		int 		`json:"jobId"`
//	TriggerTime	time.Time	`json:"triggerTime"`
//	TriggerCode	int 		`json:"triggerCode"`
//	TriggerMsg	string		`json:"triggerMsg"`
//	HandleTime	time.Time 	`json:"handleTime"`
//	HandleCode  int 		`json:"handleCode"`
//	HandleMsg	string		`json:"handleMsg"`
//}

type TaskLogs struct {
	Stime           string `json:"starttime"`
	Etime           string `json:"endtime"`
	RecordsFiltered int    `json:"recordsFiltered"`
	RecordsTotal    int    `json:"recordsTotal"`
	//Data				[]TaskLog	`json:"data"`
}

/*
httpclient；
*/
func HttpRequest(method, url, token string, data io.Reader) (*http.Response, error) {

	myCookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: myCookieJar}

	prereq, _err := http.NewRequest(method, url, data)
	if _err != nil {
		log.Println(_err)
	}
	if method == "POST" {
		prereq.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}

	var mycookie []*http.Cookie
	mycookie = append(mycookie, &http.Cookie{
		Name:  "XXL_JOB_LOGIN_IDENTITY",
		Value: token,
	})

	urlx, _ := httpurl.Parse(url)
	myCookieJar.SetCookies(urlx, mycookie)

	response, _err := client.Do(prereq)

	return response, _err

}

/*
获取app 执行器列表；多app
*/
func getExec(addr, token string) []*appInfo {

	url := fmt.Sprintf("http://%s/xxl-job-admin/jobgroup", addr)
	req, _err := HttpRequest("GET", url, token, nil)
	if _err != nil {
		log.Println("GET: ERROR", _err)
	}
	defer req.Body.Close()
	body, _err := goquery.NewDocumentFromReader(req.Body)
	if _err != nil {
		fmt.Println("Goquery pasrh ERROR: ", _err)
	}

	/*
		如果class内有空格，例如：<span class="text title">Go </span>; body.Find("table.text.title")
		Find()
			# 查找ID ；
			.查找class;
			ele.Find("h2").Find("a") //链式调用
	*/

	var appv1 []*appInfo
	body.Find("button.btn.btn-warning.btn-xs.update").Each(func(i int, selection *goquery.Selection) {
		appid, _ := selection.Attr("id")
		appname, _ := selection.Attr("appname")
		apptitle, _ := selection.Attr("title")
		appv1 = append(appv1, &appInfo{
			AppId:   appid,
			AppName: appname,
			AppDesc: apptitle,
		})
	})

	return appv1
}

/*
获取执行器所对应任务列表： 一个app --> 多任务
*/
func getJob(addr, token, jobGroup string) []*jobInfo {

	type resultData struct {
		RecordsTotal int       `json:"recordsTotal"`
		Data         []jobInfo `json:"data"`
	}

	url := fmt.Sprintf("http://%s/xxl-job-admin/jobinfo/pageList", addr)

	// 构造formdata参数
	formData := httpurl.Values{}
	formData.Set("jobGroup", jobGroup)
	formData.Set("jobDesc", "")
	formData.Set("executorHandler", "")
	formData.Set("start", "0")
	formData.Set("length", "1000")

	formDataByte := []byte(formData.Encode())
	formDataReader := bytes.NewReader(formDataByte)

	req, _err := HttpRequest("POST", url, token, formDataReader)
	if _err != nil {
		fmt.Println("Post Job ERROR: ", _err)
	}
	body, _ := ioutil.ReadAll(req.Body)

	var rData resultData
	err := json.Unmarshal(body, &rData)
	if err != nil {
		fmt.Println("Json UnMarshal ERROR:", err)
	}

	if rData.RecordsTotal == 0 {
		return nil
	}

	var job []*jobInfo
	for k, _ := range rData.Data {
		//job = append(job, &v) 此种方式错误，这样相当于值拷贝，添加的内容是&v地址所对应的值。正确方式如下：
		job = append(job, &rData.Data[k])

	}

	return job
}

/*
获取执行器执行任务结果： 一个task --> 时间范围内执行失败次数。
*/
//func GetTaskExecResult(addr, token string, taskID, status int, ch chan TaskLogs) TaskLogs {
func GetTaskExecResult(addr, token string, taskID, status int, ch chan TaskLogs) {

	const timeformat = "2006-01-02 15:04:05"
	nowTime := time.Now()
	yt := nowTime.AddDate(0, 0, -1)
	startTime := time.Date(yt.Year(), yt.Month(), yt.Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(yt.Year(), yt.Month(), yt.Day(), 23, 59, 59, 0, time.Local)

	stime := startTime.Format(timeformat)
	etime := endTime.Format(timeformat)

	filterTime := fmt.Sprintf("%s - %s", stime, etime)
	ftime := httpurl.QueryEscape(filterTime)
	url := fmt.Sprintf("http://%s/xxl-job-admin/joblog/pageList?jobGroup=0&jobId=%s&logStatus=%s&filterTime=%s", addr, strconv.Itoa(taskID), strconv.Itoa(status), ftime)

	req, _err := HttpRequest("GET", url, token, nil)
	if _err != nil {
		log.Fatal("get task result error", _err)
	}
	body, _ := ioutil.ReadAll(req.Body)

	var taskLog TaskLogs
	err := json.Unmarshal(body, &taskLog)
	if err != nil {
		log.Fatal("TaskLog Json UnMarshal ERROR:", err)
	}
	taskLog.Stime = stime
	taskLog.Etime = etime
	fmt.Println(taskID)
	ch <- taskLog
	//return taskLog
}

func ListJob(url, token string) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"应用名称", "应用描述", "JobID", "Job名称", "开始时间", "结束时间", "失败次数"})
	t.SetOutputMirror(os.Stdout)

	red := color.New(color.BgRed).SprintFunc()
	//green := color.New(color.FgGreen).SprintFunc()

	for _, app := range getExec(url, token) {
		appid := app.AppId
		appname := app.AppName
		appdesc := app.AppDesc

		jobinfo := getJob(url, token, appid)
		for _, job := range jobinfo {
			//taskSuccess := GetTaskExecResult(url, token, (*job).JobId, 1)
			//taskFailed := GetTaskExecResult(url, token, (*job).JobId, 2)
			var ch = make(chan TaskLogs, 100)
			go GetTaskExecResult(url, token, (*job).JobId, 2, ch)
			taskFailed := <-ch

			if (*job).JobStatus == "PAUSED" {
				//t.AppendRow([]interface{}{red(appname), red(appdesc), red((*job).JobId), red((*job).JobDesc), red((*job).JobStatus)})
				t.AppendRow([]interface{}{red(appname), red(appdesc), red((*job).JobId), red((*job).JobDesc), red(taskFailed.Stime), red(taskFailed.Etime), red(taskFailed.RecordsTotal)})

			} else {
				t.AppendRow([]interface{}{appname, appdesc, (*job).JobId, (*job).JobDesc, taskFailed.Stime, taskFailed.Etime, taskFailed.RecordsTotal})

			}
			//fmt.Println(appname,appdesc, (*job).JobDesc, (*job).JobStatus)
		}
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}

func HandlerJob(operate, addr, token string, jobId int) int {
	var url string
	if operate == "pause" {
		url = fmt.Sprintf("http://%s/xxl-job-admin/jobinfo/pause", addr)
	} else if operate == "resume" {
		url = fmt.Sprintf("http://%s/xxl-job-admin/jobinfo/resume", addr)
		fmt.Println(url)
	}

	// 构造formdata参数pause
	formData := httpurl.Values{}
	formData.Set("id", strconv.Itoa(jobId))

	formDataByte := []byte(formData.Encode())
	formDataReader := bytes.NewReader(formDataByte)

	req, _err := HttpRequest("POST", url, token, formDataReader)
	if _err != nil {
		fmt.Println("Job Handler ERROR: ", _err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	status_code := req.StatusCode
	fmt.Println(status_code)
	fmt.Println(string(body))
	type Data struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		Content string `json:"content"`
	}
	var rData Data
	err := json.Unmarshal(body, &rData)
	if err != nil {
		fmt.Println("Json UnMarshal ERROR:", err)
	}

	return rData.Code
}

/*
func buildYamlConfig(url, token string) {

	type Controller struct {
		PauseController		[]map[string]string		`yaml:"pauseController"`
		ResumeController	[]map[string]string		`yaml:"resumeController"`
	}

	type YamlConfig struct {
		Xxljob 	Controller	`yaml:"xxljob"`
	}

	var data []map[string]string
	for _, app := range getExec(url, token) {
		appid := app.AppId
		appname := app.AppName

		jobinfo := getJob(url, token, appid)
		for _, job := range jobinfo {
			data_v1 := map[string]string{
				"appName": appname,
				"jobDesc": (*job).JobDesc,
				"jobId": strconv.Itoa((*job).JobId),
			}
			data = append(data, data_v1)
		}
	}
	controller := Controller{
		PauseController: data,
		ResumeController: data,
	}
	yconfig := YamlConfig{
		Xxljob: controller,
	}

	result, _err := yaml.Marshal(yconfig)
	if _err != nil {
		fmt.Println("yaml Marshal Error:", _err)
	}
	fmt.Println(string(result))
}
*/

func main() {

	URL := pflag.StringP("url", "u", "", "Setting XXL-JOB-ADMIN URL. Module: -u 172.100.10.15:8080 or test-xxl.xxx.com:80")
	TOKEN := pflag.StringP("token", "t", "", "Setting XXL-JOB-ADMIN AUTH. Module: -t 'xxxxx'")
	OPERATE := pflag.StringP("operate", "o", "", "Value: resume or pause. 'resume' Recovery Service; 'pause' stop Service")
	CONFIRM := pflag.String("confirm", "", "Handler resume and pause, Confirm Parmeter. ")
	//STARTTIME := pflag.StringP("starttime","s", "", "Get xxljob task exec log. Getting Start Time. ")
	//ENDTIME := pflag.StringP("endtime","e", "","Get xxljob task exec log. Getting End Time. ")

	pflag.Parse()

	if *URL == "" {
		panic(fmt.Errorf("Please Input URL !!!"))
	}
	if *TOKEN == "" {
		panic(fmt.Errorf("Please Input Auth Token !!!"))
	}

	red := color.New(color.BgRed).SprintFunc()
	//green := color.New(color.FgGreen).SprintFunc()
	switch *OPERATE {
	case "":
		ListJob(*URL, *TOKEN)
		//buildYamlConfig(*URL, *TOKEN)
	case "pause":
		if *CONFIRM == "execconfirm" {
			for _, app := range getExec(*URL, *TOKEN) {
				appid := app.AppId
				appname := app.AppName
				appdesc := app.AppDesc

				jobinfo := getJob(*URL, *TOKEN, appid)
				for _, job := range jobinfo {
					code := HandlerJob(*OPERATE, *URL, *TOKEN, (*job).JobId)
					if code == 200 {
						fmt.Println(appname, appdesc, (*job).JobDesc, "PAUSE SUCCESS")
					} else {
						fmt.Println(red(appname, appdesc), red((*job).JobDesc), red("PAUSE FAILED"))
					}
				}
			}
		} else {
			fmt.Println("Please Input Confirm Flag")
		}

	case "resume":
		if *CONFIRM == "execconfirm" {
			for _, app := range getExec(*URL, *TOKEN) {
				appid := app.AppId
				appname := app.AppName
				appdesc := app.AppDesc

				jobinfo := getJob(*URL, *TOKEN, appid)
				for _, job := range jobinfo {
					code := HandlerJob(*OPERATE, *URL, *TOKEN, (*job).JobId)
					if code == 200 {
						fmt.Println(appname, appdesc, (*job).JobDesc, "RESUME SUCCESS")
					} else {
						fmt.Println(red(appname, appdesc), red((*job).JobDesc), red("RESUME FAILED"))
					}
				}
			}
		} else {
			fmt.Println("Please Input Confirm Flag")
		}

	default:

		ListJob(*URL, *TOKEN)
	}

}
