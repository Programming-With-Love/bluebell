package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	//创建一个指针类型
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//	请求参数有误,直接返回感应
		//	这里纪录一下日志,每一个字段都要有自己的类型,记录错误就是error
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断error是不是validator的错误,如果是返回正确错误信息
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//翻译错误,翻译成中文,前面这个removeTopStruct是去除结构体名称的函数
			"msg": removeTopStruct(err.Translate(trans)),
		})
		return
	}
	//手动对请求参数校验
	//不要相信前端啊,虽然前端也可以进行校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//打印一下这个结构体
	fmt.Println(p)

	//	2.业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//	3.返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}
func LoginHandler(c *gin.Context) {
	//	获取请求参数及校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//	请求参数有误,直接返回感应
		//	这里纪录一下日志,每一个字段都要有自己的类型,记录错误就是error
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断error是不是validator的错误,如果是返回正确错误信息
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//翻译错误,翻译成中文,前面这个removeTopStruct是去除结构体名称的函数
			"msg": removeTopStruct(err.Translate(trans)),
		})
		return
	}
	//业务逻辑处理
	if err := logic.Login(p); err != nil {
		//可以看到谁正在登录失败
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名密码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
