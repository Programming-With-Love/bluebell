package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

/*
*  @author liqiqiorz
*  @data 2020/11/6 20:57
 */

const (
	ENDPOINTSS="http://oss-cn-beijing.aliyuncs.com"
	ACCESSKEYIDSS="LTAI4FyG9N8ejRFEzGKmWR5V"
	ACCESSKEYSECRENTSS="kmmlYGgVyVtduS5mpvkmxWYKdAeDHa"
	OSSBUCKETNAMESS="edu-1014"

)
func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// 创建OSSClient实例。
	client, err := oss.New(ENDPOINTSS, ACCESSKEYIDSS, ACCESSKEYSECRENTSS)
	if err != nil {
		HandleError(err)
	}
	// 获取存储空间。
	bucketName := "oss-go14"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	// 列举文件。
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			HandleError(err)
		}
		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
		//	这个输出的并不是图片的url
		//Bucket:  baijianruoliorz/github/haha.jpg

			fmt.Println("Bucket: ", object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
}