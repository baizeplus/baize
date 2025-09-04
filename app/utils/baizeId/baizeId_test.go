package baizeId

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetID(t *testing.T) {
	err := NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	l := 10000000
	ss := make([]string, l)
	for i := 0; i < l; i++ {
		go func(ii int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ss[ii] = GetId()
		}(i)
	}
	time.Sleep(time.Minute)
	// 检查重复
	seen := make(map[string]bool)
	duplicates := 0
	for _, s := range ss {
		if seen[s] {
			duplicates++
			fmt.Printf("发现重复: %s\n", s)
		} else {
			seen[s] = true
		}
	}

	if duplicates > 0 {
		fmt.Printf("总共发现 %d 个重复项\n", duplicates)
	} else {
		fmt.Println("没有发现重复项")
	}
}
