package main

import "fmt"

func main() {
	/* 局部变量定义 */
	var a int = 100
        
        var m1 map[string]string
        m1 = make(map[string]string)
        m1["a"]="aa"
        m1["b"] ="bb"
        fmt.Println(m1)
	fmt.Printf("a 的值为 : %d\n", a)

        for k,v :=range m1{
           
        }
}
