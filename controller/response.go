package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*  @author liqiqiorz
*  @data 2020/9/19 9:19
 */
//{
//	"code": 10001//错误码
//	"msg": xx,
//	"data":{
//
//}
//}

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	re := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, re)
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	re := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, re)
}

//返回token后 在authorized里面选择 bearer token 然后带上token即可访问
