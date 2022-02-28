package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io/ioutil"
)
// 参考：https://docs.docker.com/engine/api/sdk/examples/

func main() {
	ctx := context.Background()
	// 远程连接docker. 注：默认不开启
	//os.Setenv("DOCKER_HOST", "tcp://192.168.56.211:2376")
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}


	/*
	示例一：列出所有images
	 */
	fmt.Println("---1. Images List ---")
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
		fmt.Println(image.RepoTags)
		fmt.Println(image.Size/1000/1000)
	}

	/*
	示例二： 列出所有container, 状态包括未运行和运行的。
	 */
	fmt.Println("\n\n\n")
	fmt.Println("---2. Container List ---")
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Println(container)
	}

	/*
	示例三：docker tag
	 */
	fmt.Println("\n\n\n")
	fmt.Println("---3. Container tag ---")
	sourceImage := "quay.io/thanos/thanos:v0.22.0"
	targetImage := "golang/test/thanos:v0.22.0"
	tagerr := cli.ImageTag(ctx, sourceImage, targetImage)
	if tagerr != nil {
		panic(tagerr)
	}


	/*
	示例四：根据docker name查找id
	 */
	fmt.Println("\n\n\n")
	fmt.Println("---4. Image Search ---")
	sResult, err := cli.ImageSearch(ctx, "thanos", types.ImageSearchOptions{Limit: 10})

	//err := cli.ImageRemove(ctx, "", types.ImageRemoveOptions{})  // 删除镜像
	for _, r := range sResult {
		fmt.Println(r)
	}

	/*
	示例五： 推送到远程仓库中
	 */
	fmt.Println("\n\n\n")
	fmt.Println("---5. Image push ---")
	imageName := "x.x.x.x/public/prometheus:v2.26.0"
	// 添加认证
	authConfig := types.AuthConfig{
		Username: "xxx",
		Password: "xxx",
	}
	encodeJson, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodeJson)

	imagePush, err := cli.ImagePush(ctx, imageName, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}

	//io.Copy(os.Stdout, imagePush)

	// 读出imagePush内容，方式一
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(imagePush)
	//fmt.Println(buf.String())
	//imagePush.Close()

	// 读出imagePush内容，方式二
	r, err := ioutil.ReadAll(imagePush)
	if err != nil {
		panic(err)
	}
	imagePush.Close()
	fmt.Println(string(r))

	/*
	示例六、停止容器
	 */

}

