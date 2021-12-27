package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

/*
1. 处理txt文件
2. 处理json文件
3. 处理yaml文件
 */

type Calico struct {
	ApiVersion string 			`json: "apiVersion"`
	Kind string 				`json:"kind"`
	Metadata map[string]string 	`json:metadata`
	RoleRef map[string]string 	`json:"roleRef"`

	//Subjects []map[string]string 	`json:"subjects"`
	//Subjects []*map[string]string 	`json:"subjects"`
	Subjects *[]map[string]string 	`json:"subjects"`
}

func (calico Calico) calicoJson(file string) Calico {
	/*
		把json 转为 go类型
	*/
	/*
	   {
	   	"apiVersion": "rbac.authorization.k8s.io/v1beta1",
	   	"kind": "ClusterRoleBinding",
	   	"metadata": {
	   		"name": "calico-node"
	   	},
	   	"roleRef": {
	   		"apiGroup": "rbac.authorization.k8s.io",
	   		"kind": "ClusterRole",
	   		"name": "calico-node"
	   	},
	   	"subjects": [
	   		{
	   			"kind": "ServiceAccount",
	   			"name": "calico-node",
	   			"namespace": "kube-system"
	   		}
	   	]
	   }
	*/
	f, _ := os.ReadFile(file)  // 读取文件
	j5 := json.Unmarshal(f, &calico) // 返回error ,如果error == nil 解析成功，否则错误
	if j5 != nil {
		fmt.Println("map load to json error:",j5.Error())
	}
	return calico

}



func main() {

	/*
	Marshal()：Go数据对象 -> json数据
	UnMarshal()：Json数据 -> Go数据对象

	说明：
	struct、slice、array、map都可以转换成json
	struct转换成json的时候，只有字段首字母大写的才会被转换
	map转换的时候，key必须为string
	封装的时候，如果是指针，会追踪指针指向的对象进行封装
	 */
	//runtime.Breakpoint()
	//debug.PrintStack()

	/*
	把 go 类型转为 json
	 */
	// 1. struct
	type Test1 struct{
		Id int
		Name string
		Describe string
	}
	d1 := Test1{1,"haha","hello world"}
	j1, err := json.Marshal(d1)
	if err != nil {
		fmt.Println("struct load to json error:",err.Error())
	}
	//fmt.Println("j1 var Type:",reflect.TypeOf(j1))
	fmt.Println(string(j1))

	// 2. slice
	sz1 := [...]int{1,2,3,4,5,6,7}
	sp1 := sz1[:3]
	j2, err := json.Marshal(sp1)
	if err != nil {
		fmt.Println("slice load to json error:",err.Error())
	}
	fmt.Println(sp1)
	fmt.Println("j1 var Type:",reflect.TypeOf(j2))
	fmt.Println(string(j2))

	// 3. array
	j3, err := json.Marshal(sz1)
	if err != nil {
		fmt.Println("array load to json error:",err.Error())
	}
	fmt.Println(string(j3))

	// 4. map
	m1 := make(map[string]string)
	m1["k1"] = "v1"
	m1["k2"] = "v2"
	m1["k3"] = "v3"
	fmt.Println(m1)
	j4, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("map load to json error:",err.Error())
	}
	fmt.Println(string(j4))


	type test2 struct{
		Id int			`json:"id"`
		Name string		`json:"name"`
		Describe string	`json:"describe"`
	}
	t2 := test2{1,"haha","ok haha"}
	jj1, err := json.Marshal(t2)
	fmt.Println(string(jj1))


	/*
		把json 转为 go类型
	*/
	c_path, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR:",err)
	}
	fmt.Printf("当前路径：%s \n",c_path)
	filename := "test08_config.json"
	file := fmt.Sprintf("%s/%s",c_path,filename)
	fmt.Println(file)

	// 初始化函数
	var calico01 Calico
	ca := calico01.calicoJson(file)

	fmt.Println(ca.ApiVersion)
	fmt.Println(ca.Kind)
	fmt.Println(ca.RoleRef)
	fmt.Println(ca.Metadata)
	fmt.Println(ca.Subjects)

	fmt.Println("-------------")
	for k, v := range ca.Metadata {
		fmt.Println(k,v)
	}

	fmt.Println("-------------")
	for k, v := range ca.RoleRef {
		fmt.Println(k,v)
	}

	fmt.Println("-------------")
	for _, v := range *ca.Subjects { 	// Subjects *[]map[string]string
		for k, vv := range v {
			fmt.Println(k, vv)
		}
	}

	/*
	go类型 解析为 yaml
	 */
	fmt.Println("-------------yaml ---------------")
	type Yaml_m struct {
		Sz1 []string
		M1 map[string]string
	}
	type Yaml_test struct {
		Id int
		Name string
		Age int
		Y1 *Yaml_m
	}
	a1 := []string{"c1","c2","c3","c4"}
	a2 := make(map[string]string)
	a2["k1"] = "v1"
	a2["k2"] = "v2"
	y1 := Yaml_m{a1, a2}
	yy1 := Yaml_test{1,"haha", 18, &y1}
	r1, err := yaml.Marshal(&yy1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r1))

	/*
	yaml 文件 load为 go类型
	 */
	yaml_path, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR:",err)
	}
	yamlfile := "test08_config.yaml"
	yamlpath := fmt.Sprintf("%s/%s",yaml_path,yamlfile)
	fmt.Println(yamlpath)

	var yaml_01 Yaml01
	aa1 := yaml_01.yaml_test(yamlpath)
	fmt.Println(aa1)


}


type Yaml01 struct {
	Apiversion string	`yaml:"apiVersion"`
	Kind string 		`yaml:"kind"`
	Metadata map[string]string	`yaml:"metadata"`
	Roleref map[string]string	`yaml:"roleRef"`
	Subjects []map[string]string	`yaml:"subjects"`
}
func (y Yaml01) yaml_test(file string) Yaml01{

	f, _ := os.ReadFile(file)
	err := yaml.Unmarshal(f, &y)
	if err != nil {
		fmt.Println(err)
	}
	return y
}

/*
错误：
1. yaml 错误
   yaml: line 4: found character that cannot start any token
   原因： yaml 文件中有tab制表位，删除即可。

2. yaml 错误
	第二是要使用地址去改变，如果是这样：
	v := reflect.ValueOf(person)，获取值类型去修改的话，会报错：
	panic: reflect: reflect.Value.SetString using unaddressable value
	报的没法地址到value。

	第二个错误这儿就是 在解析时参数必须引用地址。 yaml.Unmarshal(f, &y)  yaml.Marshal(&yy1)

 */