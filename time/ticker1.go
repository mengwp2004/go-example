package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {

    ticker := time.NewTicker(time.Second * 1)
    wg := sync.WaitGroup{}

    wg.Add(10)
    go func() {
        for t := range ticker.C {
            fmt.Printf("Backup at %s\n", t)
            wg.Done()
        }
    }()

    wg.Wait()
    ticker.Stop()
}
