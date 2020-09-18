package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

/*
*  @author liqiqiorz
*  @data 2020/9/18 21:17
 */

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) {
	//	判断用户存不存在
	mysql.QueryUserByUsername()
	//生成UID
	snowflake.GenID()
	//	用户密码加密
	//保存数据仓库
	mysql.InsertUser()
}
