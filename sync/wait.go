package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i > 5; i = i + 1 {
		wg.Add(1)
                go EchoNumber(i)
		//go func(n int) {
		//	// defer wg.Done(),注意这个Done的位置，是另一个函数
		//	defer wg.Done()
		//	EchoNumber(n)
		//}(i)
	}

	wg.Wait()
        fmt.Println("after wait")
}

func EchoNumber(i int) {
        //defer wg.Done()
	time.Sleep(10000)
	fmt.Println(i)
        wg.Done()
}
