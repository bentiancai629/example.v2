package main

import "fmt"

func main() {
	var ok bool

	//双向通道
	ch := make(chan int, 1)
	_, ok = interface{}(ch).(chan int)
	fmt.Println("chan int => <-chan int:",ok)
	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => chan<- int:", ok)

	//单向发送通道
	sch := make(chan<- int, 1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

	//单向接收通道
	rch := make(<-chan int, 1)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("<-chan int => chan int:", ok)
}
