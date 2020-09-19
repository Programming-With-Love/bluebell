package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

/*
*  @author liqiqiorz
*  @data 2020/9/18 21:17
 */

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//	判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//	用户密码加密
	//保存数据仓库
	return mysql.InsertUser(user)
}
func Login(p *models.ParamLogin) (token string, err error) {

	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针,可以拿到user.userID
	if err := mysql.Login(user); err != nil {
		//	登录失败
		return "", err
	}
	//	生成JWT
	return jwt.GenToken(user.UserID, user.Username)
}
