package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"gopkg.in/errgo.v2/errors"
)

/*
*  @author liqiqiorz
*  @data 2020/9/18 21:17
 */

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) {
	//	判断用户存不存在
	if mysql.CheckUserExist(p.Username) {
		return errors.New("用户已经存在")
	}
	//生成UID
	userID := snowflake.GenID()
	u := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//	用户密码加密
	//保存数据仓库
	mysql.InsertUser()
}
