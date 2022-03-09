package main

import (
	"fmt"
	"reflect"
)

type Monster2 struct {
	Name string  `json:"monsterName"`
	Age  int     `json:"monsterAge"`
	Sal  float64 `json:"monsterSal"`
	Sex  string  `json:"monsterSex"`
}

func main() {
	var num int = 100
	reflectTest(num)
	TypeJudge(num)

	ptrReflectTest(num)
	ptrReflectTest(&num)

	var monster Monster2 = Monster2{
		Name: "abc",
		Age:  16,
		Sal:  1000,
		Sex:  "male",
	}
	reflectGetFields(monster)

}

func reflectGetFields(b interface{}) {
	reflectValue := reflect.ValueOf(b)
	if reflectValue.Kind() != reflect.Struct {
		fmt.Println("data must be struct type")
	} else {
		fmt.Printf("struct has %d field\n", reflectValue.NumField())
		for i := 0; i < reflectValue.NumField(); i++ {
			tagVal := reflect.TypeOf(b).Field(i).Tag.Get("json")
			fieldVal := reflectValue.Field(i)
			fmt.Printf("tagVal = %v, FieldValue = %v\n", tagVal, fieldVal)
		}
	}
}

func ptrReflectTest(b interface{}) {
	reflectKind := reflect.ValueOf(b).Kind()
	fmt.Println("reflectKind =", reflectKind)
	fmt.Println("reflectKind =", reflectKind.String())
	fmt.Printf("reflectKind String() %T\n", reflectKind.String())

	if reflectKind.String() == "int" {
		fmt.Println("reflectValue =", reflect.ValueOf(b).Int())
	}
	if reflectKind.String() == "ptr" {
		reflect.ValueOf(b).Elem().SetInt(30)
		fmt.Println("reflectValue =", reflect.ValueOf(b).Elem().Int())
	}

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

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool 值是%v\n", i, x)
		case float64:
			fmt.Printf("param #%d is a float64 值是%v\n", i, x)
		case int, int64:
			fmt.Printf("param #%d is a int 值是%v\n", i, x)
		case string:
			fmt.Printf("param #%s is a string 值是%v\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil 值是%v\n", i, x)
		default:
			fmt.Printf("param #%d is a unkunow 值是%v\n", i, x)
		}
	}

}
