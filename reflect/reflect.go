package main

import (
  "fmt"
  "reflect"
)

type Foo struct {
  X string
  Y int
}

func main() {
  var i int = 123
  var f float32 = 1.23
  var l []string = []string{"a", "b", "c"}

  fmt.Println(reflect.TypeOf(i))    //int
  fmt.Println(reflect.TypeOf(f))    //float32
  fmt.Println(reflect.TypeOf(l))    //[]string

  var foo Foo
  fmt.Println(reflect.TypeOf(foo))    //main.Foo

}

