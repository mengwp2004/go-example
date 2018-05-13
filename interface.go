package main

import "io"
import "os"
//import "time"
import "bytes"
import "fmt"
func main(){

var w io.Writer
w = os.Stdout // OK: *os.File has Write method
w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
fmt.Println(w)
//w = time.Second // compile error: time.Duration lacks Write method

}
