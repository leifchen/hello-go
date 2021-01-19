package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  uint
	addr address
}
type address struct {
	province string
	city     string
}
type errorString struct {
	s string
}

// WalkRun 接口
type WalkRun interface {
	Walk()
	Run()
}

func main() {
	p := person{
		age:  18,
		name: "LeifChen",
		addr: address{
			province: "福建",
			city:     "福州",
		},
	}
	printString(p)
	printString(&p)
	printString(p.addr)
	p.Walk()
	p.Run()
}

// 实现 Stringer 接口
func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (p person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}
func (p person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}
