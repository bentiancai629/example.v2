package main

import "fmt"

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { // 用于演示接收操作。
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		//发送完以后关闭dataChan  不影响接收
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()


	//两个协程都通知通道2表示结束
	<-syncChan2
	<-syncChan2
}
