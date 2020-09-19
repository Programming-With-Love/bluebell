package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

/*
*  @author liqiqiorz
*  @data 2020/9/19 21:32
 */
//
func GetCommunityList() (data []*models.Community, err error) {
	//	查询数据
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&data, sqlStr); err != nil {
		//查询为空
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

//p51这里解决很经典,把参数实际话
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name,introduction,create_time from community where community_id=?"
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			//自定义的错误
			err = ErrorInvalidID
		}
	}
	return community, err

}
