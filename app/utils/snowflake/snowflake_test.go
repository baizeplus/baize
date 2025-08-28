package snowflake

import (
	"fmt"
	"testing"
	"time"
)

func TestGetId(t *testing.T) {
	for i := 0; i < 100; i++ {

		fmt.Println(GenID())
		time.Sleep(time.Millisecond)
	}
}
