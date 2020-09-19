package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

/*
*  @author liqiqiorz
*  @data 2020/9/19 21:30
 */
//这里就是查数据啦
func GetCommunityList() ([]*models.Community, error) {
	//	查所有数据并返回
	return mysql.GetCommunityList()

}
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
