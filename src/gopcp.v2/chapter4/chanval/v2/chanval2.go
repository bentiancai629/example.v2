package main

import (
	"fmt"
	"time"
)

/**
	指针演示 协程对同一个结构体的操作
 */
// Counter 代表计数器的类型。
type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() { // 用于演示发送操作。
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %d. [sender]\n", countMap)
			fmt.Printf("The count map: %s. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func (counter *Counter) String() string {
	return fmt.Sprintf("count: %d", counter.count)
}