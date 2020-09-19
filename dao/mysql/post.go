package mysql

import "time"

/*
*  @author liqiqiorz
*  @data 2020/9/20 0:15
 */
//内存对齐概念 尽量和数据库在同一位置
//数据库对应字段创建到这里main
type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
