package main

import (
	"encoding/json"
	"fmt"
)

// 通过字段标签的方式指定了序列化时json字段的名称, 底层使用反射机制实现
type Monster struct {
	Name string  `json:"monsterName"`
	Age  int     `json:"monsterAge"`
	Sal  float64 `json:"monsterSal"`
	Sex  string  `json:"monsterSex"`
}

func main() {
	m := Monster{
		Name: "ytj",
		Age:  20,
		Sal:  888.99,
		Sex:  "female",
	}
	data, _ := json.Marshal(m)
	fmt.Println("json result:", string(data))
}
