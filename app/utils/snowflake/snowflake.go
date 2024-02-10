package snowflake

import (
	"baize/app/setting"
	"github.com/gogf/gf/v2/util/gconv"

	"time"

	sf "github.com/bwmarrin/snowflake"
)

var sfNode *sf.Node

func init() {
	id := 1
	var st time.Time
	st, err := time.Parse("2006-01-02", setting.Conf.StartTime)
	if err != nil {
		panic(err)
	}
	sf.Epoch = st.UnixNano() / 1000000
	sfNode, err = sf.NewNode(gconv.Int64(id))
	if err != nil {
		panic(err)
	}
}

func GenID() int64 {
	return sfNode.Generate().Int64()
}
