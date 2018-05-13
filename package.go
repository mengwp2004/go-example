package  main

import "fmt"

func main(){

  var a = b + c // a 第三个初始化, 为 3
  var b = f() // b 第二个初始化, 为 2, 通过调用 f (依赖c)
  var c = 1 // c 第一个初始化, 为 1

  func f() int{
     return c + 1
  }

  fmt.Println(a)
｝
