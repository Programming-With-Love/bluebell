package controller

import "github.com/gin-gonic/gin"

/*
*  @author liqiqiorz
*  @data 2020/9/20 0:12
 */
func CreatePostHandler(c *gin.Context) {
	//1.获取参数以及参数的校验
	c.ShouldBindJSON()
	//	2.创建帖子
	//	返回相应
}
