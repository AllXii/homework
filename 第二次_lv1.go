package main

import "fmt"

type person struct {
	name string
	age int
	gender string
}
type dove interface {
	gugugu()
}
type repeater interface {
	repeat(string)
}
type sourGraper interface {
    sourGrapes()
}
type jingZe interface {
	zhenXiang()
}

func (p *person) gugugu(){
	fmt.Println(p.name + "又咕了")
}
func (p *person) repeat(words string){
	fmt.Println(words)
}
func (p *person) sourGrapes() {
	fmt.Println(p.name + "酸了")
}
func (p *person) zhenXiang(){
	fmt.Println(p.name + ":\"真香!\"")
}

func main(){
    per := person{
    	name: "张三",
    	age: 18,
    	gender: "男",
	}
	per.gugugu()
    per.repeat("hello,world!")
    per.sourGrapes()
    per.zhenXiang()
}