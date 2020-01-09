package main

import (
	"fmt"
)

func Count(ch chan int) {
	fmt.Println("Counting ", ch)
	ch <- 2
}

func main() {
	//var m map[string] chan bool
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
	//for i := 0; i < 10; i++ {
	//	go Add(i, i)
	//}
	//time.Sleep(2000)
	println("bfbctest1.main")
}

func Add(i int, j int) {
	z := i + j
	println(i, "z=", z)
}
