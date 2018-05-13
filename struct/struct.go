package main

import "fmt"

func main() {
        type Person struct {
           name string
           age int
        }
        p := Person{"James", 23}  //有序
        p1 := Person{age:23,name:"mwp"}       //无序
	fmt.Printf(p.name + "\n")
        fmt.Printf(p1.name + "\n")
}
