package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

/*
*  @author liqiqiorz
*  @data 2020/9/18 21:14
 */
//SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//	1.参数校验获取参数
	//	gin框架几种携带请求参数的方式
	//	1.c.Query()
	//	2.c.params  路径参数
	//	3.ShouldBindJSON() JSON数据,必须建立在一个已有的结构体上
	//	在这里,我们用models里面的params进行了使用
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		//	请求参数有误,直接返回感应
		//	这里纪录一下日志,每一个字段都要有自己的类型,记录错误就是error
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//打印一下这个结构体
	fmt.Println(p)

	//	2.业务处理
	logic.SignUp()
	//	3.返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}
