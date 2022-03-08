package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 100
	reflectTest(num)
}

func reflectTest(b interface{}) {

	reflectType := reflect.TypeOf(b)
	fmt.Println("reflectType =", reflectType)

	// reflect.ValueOf 获取的值类型是 reflect.Value, 不再是原始的类型，如果需要对值进行操作，先要获取到原始的值对象
	reflectValue := reflect.ValueOf(b)
	fmt.Println("reflectValue =", reflectValue)
	fmt.Printf("reflectValueType = %T\n", reflectValue)

	// 获取原始的值对象并操作
	originalValue := reflectValue.Int()
	fmt.Println("originalValue =", originalValue)
	fmt.Printf("originalValueType = %T\n", originalValue)

	// 将reflect.Value重新转换为原始类型
	originalNum := reflectValue.Interface().(int)
	fmt.Println("originalNum =", originalNum)
	fmt.Printf("originalNumType = %T\n", originalNum)

}
