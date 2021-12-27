package main

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)


func main() {
	checkdir := "data-volume"

	/*
	StringSlice和 StringSliceP 区别： StringSliceP函数有简写
	 */
	home := os.Getenv("HOME")
	if home == "" {
		panic(fmt.Errorf("Please setting Config Path"))
	}
	kubeconfig := pflag.StringP("config", "c", filepath.Join(home, ".kube", "config") , "K8S AUTH CONFIG FILE")
	namespace := pflag.StringP("namespace", "n", "default", "K8S NAMESPACE")

	checkDir := pflag.StringSliceP("upload", "d", []string{}, "CHECK UPLOAD DIR LIST *Required")
	noservice := pflag.StringSliceP("noservice", "s", []string{}, "EXCEPT SERVICE LIST")

	pflag.Parse()
	if checkDir == nil {
		panic(fmt.Errorf("Please Input Check Dir"))
	}
	if noservice == nil {}

	fmt.Println("config: ", *kubeconfig)


	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	/*
	clientset.AppV1  主要是控制deployment/daemonset/RS/statusfulset等资源
	clientset.CoreV1 主要控制namespace/pod/configmap/service/secret等资源
	 */
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// List(ctx context.Context, opts metav1.ListOptions)
	deploylist, err := clientset.AppsV1().Deployments(*namespace).List(context.TODO(), metav1.ListOptions{})


	for _, value := range deploylist.Items {
		deployname := value.ObjectMeta.Name
		podVolumes := value.Spec.Template.Spec.Volumes
		containerTemp := value.Spec.Template.Spec.Containers

		fmt.Println("--------------------------------:01")
		for _, v := range podVolumes {
			if v.HostPath != nil {
				if v.Name == checkdir {
					fmt.Println(deployname, v.Name, v.HostPath.Path)
				}
			}
		}
		fmt.Println("--------------------------------:02")
		for _, v := range containerTemp {
			for _, v1 := range v.VolumeMounts {
				if v1.Name == checkdir {
					fmt.Println(v1.Name,v1.MountPath)
				}
			}
		}
	}
}

