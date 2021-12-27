package main

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"github.com/jedib0t/go-pretty/table"
)


func main() {

	kubeconfig := pflag.StringP("config", "c", filepath.Join(getHome(), ".kube", "config") , "K8S AUTH CONFIG FILE")
	namespace := pflag.StringP("namespace", "n", "default", "K8S NAMESPACE")

	checkPath := pflag.StringSliceP("checkpath", "p", nil, "CHECK UPLOAD DIR LIST *Required")
	service := pflag.StringSliceP("service", "s", nil, "CHECK SERVICE LIST")
	exceptservice := pflag.StringSliceP("except-service", "e", nil, "")
	// 使用 --noservice="prometheus,test"
	pflag.Parse()

	//if *checkPath == nil {
	//	panic(fmt.Errorf("checkpath args cannot empty"))
	//}

	if *service != nil && *exceptservice != nil {
		panic(fmt.Errorf("service AND exceptservice args cannot exist at the same time!"))
	}

	if *service != nil || *exceptservice != nil {
		selectService(kubeconfig, namespace, checkPath, service, exceptservice)
	}

}

func selectService(kubeconfig, namespace *string, checkPath, service, exceptservice *[]string) {

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Deployment Name", "volumePath", "MountPath"})
	t.SetOutputMirror(os.Stdout)
	if *service != nil {
		for _, value := range *service {
			deploy, err := getDeployment(kubeconfig, namespace, strings.TrimSpace(value))
			if err != nil {
				fmt.Printf("deployment %s not exist \n", value)
				continue
				//panic(err.Error())
			}
			for _, value1 := range *checkPath {
				dname, vp, mp := getVolumeInfo(strings.TrimSpace(value1),deploy)

				if vp == "" && mp == "" {
					//fmt.Printf("%s: not mount %s \n", dname, value1)
					t.AppendRow([]interface{}{dname,  vp, mp})
				} else {
					//fmt.Printf("servicename: %s , hostpath: %s, mountpath: %s \n",dname, vp, mp)
					t.AppendRow([]interface{}{dname,  vp, mp})
				}
			}
		}
	}

	if *exceptservice != nil {
		deploylist, err := getAllDeployments(kubeconfig, namespace)
		if err != nil {
			panic(fmt.Errorf("Namespace %s cannot exist services. ERROR: %s", namespace, err))
		}
		for _, value := range deploylist.Items {

			deployName := value.Name

			// 判断 value 是否在  exceptservice 中
			sort.Strings(*exceptservice)
			i := sort.SearchStrings(*exceptservice, deployName)
			if i < len(*exceptservice) && (*exceptservice)[i] == deployName {
				continue
			}

			for _, value1 := range *checkPath {
				deployname, vp, mp := getVolumeInfo(value1, &value)
				if vp == "" && mp == "" {
					//fmt.Printf("%s: not mount %s \n", deployname, value1)
					t.AppendRow([]interface{}{deployname,  vp, mp})
				} else {
					//fmt.Printf("servicename: %s , hostpath: %s, mountpath: %s \n",deployname, vp, mp)
					t.AppendRow([]interface{}{deployname,  vp, mp})
				}

			}

		}

	}
	t.Render()


}

func getVolumeInfo(vname string, deployment *v1.Deployment) (deployName, volumePath, mountPath string) {

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


func getDeployment(kubeconfig *string, namespace *string, service string) (*v1.Deployment, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	serviceDeployment, err := clientset.AppsV1().Deployments(*namespace).Get(context.TODO(), service, metav1.GetOptions{})
	return serviceDeployment, err
}

func getAllDeployments(kubeconfig *string, namespace *string) (*v1.DeploymentList, error) {

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

func getHome() string {
	home := os.Getenv("HOME")
	if home == "" {
		panic(fmt.Errorf("Please setting Config Path"))
	}
	return home
}

