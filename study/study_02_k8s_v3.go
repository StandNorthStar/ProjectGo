package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/pflag"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"sort"
)

func intersect(s1 []string, s2 []string) []string {
	m := make(map[string]int)
	for _,v := range s1 {
		m[v]++
	}
	fmt.Println(m)
	for _,v := range s2 {
		times, _ := m[v]  //v是nums2中的值,m[v]是map中的值.m[v]==times
		if times == 0{
			s1 = append(s1, v)
		}
	}
	return s1
}

func getHomeV3() string {
	home := os.Getenv("HOME")
	if home == "" {
		panic(fmt.Errorf("Please setting Config Path"))
	}
	return home
}

func getAllDeploymentsV3(kubeconfig *string, namespace *string) (*v1.DeploymentList, error) {

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deploylist, err := clientset.AppsV1().Deployments(*namespace).List(context.TODO(), metav1.ListOptions{})
	return deploylist, err
}

func getVolumeInfoV3(vname string, deployment *v1.Deployment) (volumePath, mountPath string) {
	/*
	根据 volumepath 找到 name ；
	然后根据name找到 mountpath  ;
	   如果可以找到mountpath 说明挂载成功； 找不到mountpath 挂载失败
 	*/
	var volumepath string
	var mountpath string
	podVolumes := deployment.Spec.Template.Spec.Volumes
	container := (deployment.Spec.Template.Spec.Containers)[0]
	// 遍历pod的Volume列表
	for _, value := range podVolumes {

		// 如果没有挂载 HostPath是nil
		if value.HostPath == nil {
			continue
		}

		// 判断vname目录是否在Volume列表中。
		if value.HostPath.Path == vname {
			volumepath = value.HostPath.Path
			hostpathname := value.Name

			// 遍历pod的Mount列表
			for _, mpath := range container.VolumeMounts {
				// 判断vname对应目录的VolumeName和MountName是否一样
				if mpath.Name == hostpathname {
					mountpath = mpath.MountPath
					return volumepath, mountpath
				}
			}

		}
	}

	return volumepath, mountpath
}



func checkPath(kubeconfig, namespace *string, checkpath []string) {
	/*
	思路：
	1. 从命令行参数传递进来目录；
	2. 便利所有deployment服务，查出deploy名称、vloume名称--> mount目录；
	3. 判断 deployment是否在 path中；
		如果在，从path中读取deployment名称对应的value值；
			便利参数列表，判断参数是否在value值里面；
				如果在， 输出加*； 如果不在正常输出。
		如果不在，正常获取内容。
	 */
	path := make(map[string][]string)
	path["xxx-xxx-xxx"] = []string{"/xxx", "/xxx"}
	path["xxx-xxx-xxx"] = []string{"/xxx", "/xxx"}



	deploylist, err := getAllDeploymentsV3(kubeconfig, namespace)
	if err != nil {
		panic(fmt.Errorf("Namespace %s cannot exist services. ERROR: %s", namespace, err))
	}


	red := color.New(color.BgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	t := table.NewWriter()
	t.AppendHeader(table.Row{"应用列表", "检测目录", "volume目录", "Pod挂载目录", "必挂载项","挂载状态"})
	t.SetOutputMirror(os.Stdout)  // 控制输出控制台或者文件 io.Writer

	for _, value := range deploylist.Items {
		deployname := value.Name
		pValue, judge := path[deployname]
		if judge == false {
			//fmt.Println("This Deployment Not In List")
			if checkpath == nil {
				continue
			} else {
				for _, argsPath := range checkpath {
					//vp, mp := getVolumeInfoV2(argsPath, &value)
					volumepath, mountpath := getVolumeInfoV3(argsPath, &value)
					if volumepath != "" && mountpath != "" {
						t.AppendRow([]interface{}{deployname, argsPath, volumepath, mountpath, "", "OK"})
					} else {
						t.AppendRow([]interface{}{deployname, argsPath, volumepath, mountpath, "", "ERROR"})
					}
				}
			}
		} else {

			// 合并必挂目录和参数指定目录
			newPathList := []string{}
			if len(checkpath) == 0 {
				newPathList = append(newPathList, pValue...)
			} else {
				for _, argsPath := range checkpath {
					// 1. 先把获取到的pValue添加到 newPathList
					newPathList = intersect(newPathList, pValue)

					// 2. 判断参数checkpath目录是否在pValue对应的目录中
					sort.Strings(pValue)
					i := sort.SearchStrings(pValue, argsPath)
					if i < len(pValue) && (pValue)[i] == argsPath {
						//fmt.Printf("%s path has exsit \n", argsPath)
					} else {
						newPathList = append(newPathList, argsPath)
					}
				}
			}
			//fmt.Println(newPathList)

			for _, newArgsPath := range newPathList {
				volumepath, mountpath := getVolumeInfoV3(newArgsPath, &value)

				sort.Strings(pValue)
				i := sort.SearchStrings(pValue, newArgsPath)
				//fmt.Println(i < len(pValue) && (pValue)[i] == newArgsPath)
				if i < len(pValue) && (pValue)[i] == newArgsPath {
					if volumepath != "" && mountpath != "" {
						t.AppendRow([]interface{}{green(deployname), green(newArgsPath), green(volumepath), green(mountpath), green("*"), green("OK")})
					} else {
						t.AppendRow([]interface{}{red(deployname), red(newArgsPath), red(volumepath), red(mountpath), red("*"), red("ERROR")})
					}
				} else {
					if volumepath != "" && mountpath != "" {
						t.AppendRow([]interface{}{deployname, newArgsPath, volumepath, mountpath, "", "OK"})
					} else {
						t.AppendRow([]interface{}{deployname, newArgsPath, volumepath, mountpath, "", "ERROR"})
					}

				}

			}
		}

	}
	//t.SetStyle(table.StyleColoredBright)
	t.SetStyle(table.StyleLight)
	t.Render()

}


func main() {

	kubeconfig := pflag.StringP("config", "c", filepath.Join(getHomeV3(), ".kube", "config"), "K8S AUTH CONFIG FILE")
	namespace := pflag.StringP("namespace", "n", "default", "K8S NAMESPACE")
	CHECKPATH := pflag.StringSliceP("checkpath", "p", nil, "CHECK UPLOAD DIR LIST *Required")
	//service := pflag.StringSliceP("service", "s", nil, "CHECK SERVICE LIST")
	// 使用 --noservice="prometheus,test"
	pflag.Parse()


	if *kubeconfig == "" {
		*kubeconfig = filepath.Join(getHomeV3(), ".kube", "config")
	}
	if *namespace == "" {
		*namespace = "default"
	}
	// 目录处理
	if CHECKPATH != nil {

	}
	checkPath(kubeconfig, namespace, *CHECKPATH)

}



