package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

//调用直接生成ID值
var node *sf.Node

//传过来上线时间 ,整个分布系统有关表示
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
