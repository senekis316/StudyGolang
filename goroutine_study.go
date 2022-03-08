package main

import (
	"fmt"
)

func writeChannel(intChannel chan int) {
	for i := 1; i <= 50; i++ {
		fmt.Println("writeValue =", i)
		intChannel <- i
	}
	close(intChannel)
}

func readChannel(intChannel chan int, extChannel chan bool) {
	for {
		v, ok := <-intChannel
		if !ok {
			break
		}
		fmt.Println("readValue =", v)
	}
	extChannel <- true
	close(extChannel)
}

func main() {

	intChannel := make(chan int, 50)
	extChannel := make(chan bool, 1)

	go writeChannel(intChannel)
	go readChannel(intChannel, extChannel)

	for {
		v, ok := <-extChannel
		if !ok {
			break
		}
		fmt.Println("读取推出信号:", v)
	}
	fmt.Println("数据读取完成, 程序退出")

}
