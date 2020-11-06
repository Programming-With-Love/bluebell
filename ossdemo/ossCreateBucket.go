package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

/*
*  @author liqiqiorz
*  @data 2020/11/6 20:08
 */

const (
	ENDPOINT="http://oss-cn-beijing.aliyuncs.com"
	ACCESSKEYID="LTAI4FyG9N8ejRFEzGKmWR5V"
	ACCESSKEYSECRENT="kmmlYGgVyVtduS5mpvkmxWYKdAeDHa"
	OSSBUCKETNAME="edu-1014"

)
func handleError(err error){
	fmt.Println("Error",err)
	os.Exit(-1)
}
//以下程序可以创建一个bucket示例
func main() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	 endpoint := ENDPOINT
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := ACCESSKEYID
	accessKeySecret := ACCESSKEYSECRENT

	bucketName := "oss-go14"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 创建存储空间。
	err = client.CreateBucket(bucketName)
	if err != nil {
		handleError(err)
	}
}