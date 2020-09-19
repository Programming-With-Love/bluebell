package models

import "time"

/*
*  @author liqiqiorz
*  @data 2020/9/19 21:33
 */
type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
type CommunityDetail struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
	//omitempty表示如果是空的话就不展示了
	Introduction string `json:"introduction,omitempty",db:"introduction"`
	//使用time.time类型的时候,数据库用的是时间戳,所以连接mysql要加上parseTime=true
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
