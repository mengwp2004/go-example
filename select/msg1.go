package main

import (
    "time"
    "log"
    "strconv"
)


func hello(){
   log.Println("hello(),frieds!");
}
func receiveMsg(){
    c1 := make(chan string)
    c2 := make(chan string)
    t1 := time.Now() // get current time
    go func(){
        //received three message
        for i:=0; i < 3; i++ {
            time.Sleep(time.Millisecond * 150)
            //c1 <- "msg 1 with index " + strconv.Itoa(i)
        }
    }()
    go func(){
        for j:=0; j < 3; j++ {
            time.Sleep(time.Millisecond * 200)
            c2 <- "msg 2 with index "+ strconv.Itoa(j)
        }
        <- c1
    }()

    //print two message
    for i:= 0; i < 4; i++ {
        msg1:="hello"
        select {
        case  c1 <- msg1:
            log.Println("received msg :", msg1)
            hello()
        case msg2 := <- c2:
            log.Println("received msg :", msg2)
        }
    }
    elapsed := time.Since(t1)
    log.Println("time elapsed ", elapsed)

}
func main() {
    receiveMsg()

}
