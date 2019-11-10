package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	//两个通道
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	//协程1
	go func() { // 用于演示接收操作。
		<-syncChan1 //接收通道信号 等待非空 被唤醒
		fmt.Println("Received a sync signal and wait a second... [receiver]")

		//睡眠1s
		time.Sleep(time.Second)

		//从通道循环接收数据
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")

		//唤醒通道2
		syncChan2 <- struct{}{}  //不包含任何字段的结构体类型 不占用内存 所有该类型变量都共享内存
	}()

	//协程2
	go func() { // 用于演示发送操作。
		for id, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", "id-",id,":",elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 seconds... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	//接收到2此信号 主函数协程结束
	<-syncChan2
	<-syncChan2
}
