package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

//因为多个main 所以必须重新定义这些东西和函数,但是保持原名会报错
const (
	ENDPOINTS="http://oss-cn-beijing.aliyuncs.com"
	ACCESSKEYIDS="LTAI4FyG9N8ejRFEzGKmWR5V"
	ACCESSKEYSECRENTS="kmmlYGgVyVtduS5mpvkmxWYKdAeDHa"
	OSSBUCKETNAMES="edu-1014"

)
/*
*  @author liqiqiorz
*  @data 2020/11/6 20:49
 */
func handleError1(err error){
	fmt.Println("Error",err)
	os.Exit(-1)
}
func main() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := ENDPOINTS
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := ACCESSKEYIDS
	accessKeySecret := ACCESSKEYSECRENTS
	bucketName := "oss-go14"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。

	//这里是上传的jpg的名称 会自动创建问价加
	objectName := "baijianruoliorz/github/haha.jpg"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	//实体名称
	localFileName := "D:\\JPG\\Ekc7xAgU0AAWzKz.jpg"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError1(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError1(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError1(err)
	}
}