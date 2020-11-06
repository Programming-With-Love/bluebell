package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

/*
*  @author liqiqiorz
*  @data 2020/11/6 20:01
 */

func main(){
	getString := viper.GetInt("redis.port")

	fmt.Println("OSS end point is :")
	fmt.Println("OSS end point is :",getString)

	fmt.Println("OSS GO SDK version: ",oss.Version)

}