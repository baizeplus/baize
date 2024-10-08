package task

import (
	"fmt"
)

func NoParams(p ...string) {
	fmt.Println("无参测试")
}

func Params(p ...string) {
	if len(p) != 2 {
		fmt.Println("参数错误")
	}
	fmt.Printf("有参数测试:\t名字 = %s,\t年龄 = %d\n", p[0], p[1])
}
