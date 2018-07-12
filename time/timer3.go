package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan struct{}) // 用来等待协程结束

    go time.AfterFunc(time.Second*3, func() {
        fmt.Println("Print after 3 seconds")
        done <- struct{}{} // 通知主协程结束
    })

    fmt.Println("Print in main")

    <-done
    close(done)
}
