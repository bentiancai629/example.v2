package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}

	//顺序遍历
	//for index,b := range names{
	//	fmt.Println(index,b)
	//}

	//并发遍历
	for id, names := range names {
		go func(id int,who string) {
			fmt.Printf("hello,%v, %s!\n",id,who)
		}(id,names)
	}

	time.Sleep(time.Microsecond * 10)

	//for _, name := range names {
	//	go func(who string) {
	//		fmt.Printf("Hello, %s!\n", who)
	//	}(name)
	//}
	//time.Sleep(time.Millisecond)
}
