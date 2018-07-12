package main

import (
   "time"
   "fmt"
)


func main(){

   pingLoop()
}

func  pingLoop() {
        ping := time.NewTimer(2*time.Second)
        for {
                select {
                case <-ping.C:
                     fmt.Println("ping")
                }
        }
}

