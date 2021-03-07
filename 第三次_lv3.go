package main

import (
	"fmt"
	"time"
)

//@myContext  罗振玺
//@初始化方法   v := make(myContext)
//@通知关闭使用方法cancel()
//v是一个管道,主函数调用v.cancel()时就会关闭管道
type myContext chan struct{}

func (myCtx myContext) cancel(){
	close(myCtx)
}

func worker(t *time.Ticker, myCtx myContext){
	go func() {
		defer fmt.Println("程序结束")
		for {
			select {
			case <- myCtx:
				fmt.Println("recv stop signal")
				return
			case <-t.C:
				fmt.Println("is working")
			}
		}
	}()
}

func main(){
	ticker := time.NewTicker(time.Second)
	myCtx := make(myContext)
	worker(ticker, myCtx)
	time.Sleep(2 * time.Second)
	myCtx.cancel()
	time.Sleep(1*time.Second)
}