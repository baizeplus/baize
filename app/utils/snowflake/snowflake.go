package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

var (
	sfNode *sf.Node
	encode = []byte("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
)

func init() {
	var id int64 = 1
	var st time.Time
	st, err := time.Parse("2006-01-02", viper.GetString("start_time"))
	//st, err := time.Parse("2006-01-02", "2025-01-01")
	if err != nil {
		panic(err)
	}
	sf.Epoch = st.UnixNano() / 1000000
	sfNode, err = sf.NewNode(id)
	if err != nil {
		panic(err)
	}
}

func GenID() string {
	return toBase32(sfNode.Generate().Int64())
}

func toBase32(f int64) string {
	if f < 32 {
		return string(encode[f])
	}

	b := make([]byte, 0, 12)
	for f >= 32 {
		b = append(b, encode[f%32])
		f /= 32
	}
	b = append(b, encode[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}
