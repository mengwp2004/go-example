package main

import (
    "log"
    "fmt"
)


func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func exeFibonacci(){
    c := make(chan int)
    quit := make(chan int)

    //produce data
    go func(){
        for i:= 0; i < 10; i++ {
            log.Println("channel data item ", <- c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}

func main() {
    exeFibonacci()
}
