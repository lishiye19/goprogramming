package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Count(ch chan int) {
	fmt.Println("Counting ", ch)
	ch <- 2
}

var one sync.Once

func onehs() {
	fmt.Println("hello world")
}

func doone() {

	one.Do(onehs)

}

func main() {
	one.Do(onehs)
	one.Do(onehs)
	one.Do(onehs)
	runtime.GOMAXPROCS(8)
	icpu := runtime.NumCPU()
	fmt.Println("icpu=", icpu)

	//ch := make(chan int, 1)
	//	//for {
	//	//	select {
	//	//	case ch <- 1:
	//	//	case ch <- 2:
	//	//	}
	//	//	i := <-ch
	//	//	fmt.Println("value recived:", i)
	//	//
	//	//}
	//c := make(chan int, 1024)
	//for i := range c {
	//	fmt.Println("recived:", i)
	//}

	//var m map[string] chan bool
	//chs := make([]chan int, 10)
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go Count(chs[i])
	//}
	//for _, ch := range chs {
	//	<-ch
	//}
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
