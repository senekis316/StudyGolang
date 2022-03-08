package main

import (
	"fmt"
)

func main() {
	// defer + recover来捕获下方map抛出的错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	// 因为未使用make来初始化map对象而直接使用，导致报错
	var myMap map[int]string
	myMap[0] = "golang"
}
