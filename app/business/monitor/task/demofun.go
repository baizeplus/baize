package task

import (
	"fmt"
	"strconv"
)

func NoParams(p ...string) {
	fmt.Println("无参测试")
}

func Params(p ...string) {
	if len(p) != 2 {
		fmt.Println("参数错误")
	}
	atoi, err := strconv.Atoi(p[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("有参数测试:\t名字 = %s,\t年龄 = %d\n", p[0], atoi)
}
