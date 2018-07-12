package main

import (
    "io"
    "log"
    "net"
)

func main() {
    l, err := net.Listen("tcp", ":2000")
    checkErr(err)

    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
            continue
        }

        go func(c net.Conn) {
            io.Copy(c, c)
            c.Close()
        }(conn)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
