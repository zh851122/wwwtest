package main

import (
	"fmt"
	"time"
)

func Receive(ch chan int) {
	//先等几秒后再接收消息
	time.Sleep(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			//接收信道里边的消息，接收后缓冲就充足了

			//信道被关闭了，退出
			if !ok {
				fmt.Println("chan clos,receive:", v)
			}
			fmt.Println("receive:", v)
		}
	}
}
func Send(ch chan int) {
	//发到第11个时,会卡住，因为信道满了
	for i := 0; i < 13; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
	//打印完毕，关闭信道
	close(ch)
}
func main() {
	//新建一个5个缓冲的信道
	ch := make(chan int, 10)
	//将信道传入函数，开启协程
	go Receive(ch)
	go Send(ch)
	for {
		time.Sleep(1 * time.Second)
	}
}
