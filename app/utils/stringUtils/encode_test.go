package stringUtils

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		uuid := GetUUID()
		fmt.Println(uuid)
	}
	uuid := GetUUID()
	fmt.Println(uuid)
}
