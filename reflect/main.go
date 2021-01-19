package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// 反射：reflect.Value 和 reflect.Type
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it)

	// 修改 struct 结构体字段的值
	p := person{Name: "LeifChen", Age: 20}
	ppv := reflect.ValueOf(&p)
	fmt.Println(ppv.Kind())
	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())
	ppv.Elem().Field(0).SetString("Modified")
	fmt.Println(p)

	//struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	//json to struct
	respJSON := "{\"name\":\"李四\",\"age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)
}

type person struct {
	Name string
	Age  int
}
