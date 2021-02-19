package main

import (
	"fmt"
	"reflect"
)
func Receiver (v interface{}){
	fmt.Print("这个是")
	switch v.(type) {
	case int:
		fmt.Println(reflect.TypeOf(v))
	case string:
		fmt.Println(reflect.TypeOf(v))
	case float64:
		fmt.Println(reflect.TypeOf(v))
	case bool:
		fmt.Println(reflect.TypeOf(v))
	case byte:
		fmt.Println(reflect.TypeOf(v))
	case []int:
		fmt.Println(reflect.TypeOf(v))
	}
}
func main(){

    Receiver("你好吗")
    Receiver(32)
    Receiver(true)
    //fmt.Println(reflect.TypeOf())
}
