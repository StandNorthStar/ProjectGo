package main

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/pflag"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"sort"
)

func getHomeV2() string {
	home := os.Getenv("HOME")
	if home == "" {
		panic(fmt.Errorf("Please setting Config Path"))
	}
	return home
}

func getAllDeploymentsV2(kubeconfig *string, namespace *string) (*v1.DeploymentList, error) {

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

func getVolumeInfoV2(vname string, deployment *v1.Deployment) (deployName, volumePath, mountPath string) {

	var vp string
	var mp string

	deployname := deployment.ObjectMeta.Name
	podVolumes := deployment.Spec.Template.Spec.Volumes
	containerTemp := deployment.Spec.Template.Spec.Containers

	for _, value := range podVolumes {
		if value.HostPath != nil {
			if value.HostPath.Path == vname {
				vp = value.HostPath.Path
				break
			} else {
				vp = ""
			}
		}
	}

	for _, value := range containerTemp {
		for _, value1 := range value.VolumeMounts {
			if value1.MountPath == vname {
				mp = value1.MountPath
				break
			} else {
				mp = ""
			}
		}
	}

	return deployname, vp, mp
}



func main() {

	var checkpath []string
	checkpath = []string{"/cupload1", "/cupload"}


	kubeconfig := pflag.StringP("config", "c", filepath.Join(getHomeV2(), ".kube", "config"), "K8S AUTH CONFIG FILE")
	namespace := pflag.StringP("namespace", "n", "default", "K8S NAMESPACE")
	CHECKPATH := pflag.StringSliceP("checkpath", "p", nil, "CHECK UPLOAD DIR LIST *Required")
	//service := pflag.StringSliceP("service", "s", nil, "CHECK SERVICE LIST")
	// 使用 --noservice="prometheus,test"
	pflag.Parse()


	if *kubeconfig == "" {
		*kubeconfig = filepath.Join(getHomeV2(), ".kube", "config")
	}
	if *namespace == "" {
		*namespace = "default"
	}
	// 目录处理
	if CHECKPATH != nil {
		for _, v := range *CHECKPATH {
			sort.Strings(checkpath)
			i := sort.SearchStrings(checkpath, v)
			if i < len(checkpath) && (checkpath)[i] == v {
				continue
			} else {
				checkpath = append(checkpath, v)
			}
		}
	}

	selectServiceV2(kubeconfig, namespace, checkpath)
}

func selectServiceV2(kubeconfig, namespace *string, checkpath []string) {

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	var defaultService []string
	defaultService = []string{
		"xxx",
		"xxx",

	}

	t := table.NewWriter()
	//t.AppendHeader(table.Row{"Deployment Name", "volumePath", "MountPath", "mounted"})
	t.AppendHeader(table.Row{"应用列表", "检测目录", "volume目录", "Pod挂载目录", "必挂载项","挂载状态"})
	t.SetOutputMirror(os.Stdout)  // 控制输出控制台或者文件 io.Writer

	deploylist, err := getAllDeploymentsV2(kubeconfig, namespace)
	if err != nil {
		panic(fmt.Errorf("Namespace %s cannot exist services. ERROR: %s", namespace, err))
	}
	for _, value := range deploylist.Items {

		for _, value1 := range checkpath {
			deployname, vp, mp := getVolumeInfoV2(value1, &value)
			sort.Strings(defaultService)
			i := sort.SearchStrings(defaultService, deployname)
			if i < len(defaultService) && (defaultService)[i] == deployname {
				// 在->S中
				if vp == "" && mp == "" {
					// 没有找到指定的挂载目录 ， 标记红色
					t.AppendRow([]interface{}{red(deployname), red(value1), red(vp), red(mp), red("*"), red("ERROR")})
				} else {
					// 找到挂载目录， 标记绿色
					t.AppendRow([]interface{}{green(deployname), green(value1), green(vp), green(mp), green("*"), green("SUCESS")})
				}

			} else {
				// 不在->S中
				if vp == "" && mp == "" {
					t.AppendRow([]interface{}{deployname, value1, vp, mp, "", "-"})
					// 没有找到指定的挂载目录， 标记 -
				} else {
					t.AppendRow([]interface{}{deployname, value1, vp, mp, "", "SUCESS"})
					// 找到挂载目录， 正常输出

				}
			}
		}

	}
	//t.SetStyle(table.StyleColoredBright)
	t.SetStyle(table.StyleLight)
	t.Render()
}





