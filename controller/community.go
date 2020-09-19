package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

/*
*  @author liqiqiorz
*  @data 2020/9/19 21:24
 */
//社区相关的
func CommunityHandler(c *gin.Context) {
	//	查询到所有的社区  community_id community_name
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		//bu'qing'yi把服务端报错报给用户
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//获取社区ID,与冒号后面对应上
	communityID := c.Param("id")
	id, err := strconv.ParseInt(communityID, 10, 64)
	//如果不是字符类型,则报类型错误
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	//	查询到所有的社区  community_id community_name
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		//bu'qing'yi把服务端报错报给用户
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
