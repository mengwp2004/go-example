package main

import "fmt"

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}


var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"
fmt.Println(q)

q1 := [...]int{1, 2, 3}
fmt.Printf("%T\n", q1) // "[3]int"


type Currency int
const (
USD Currency = iota // 美元
EUR // 欧元
GBP // 英镑
RMB // 人民币
)

symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
fmt.Println(RMB, symbol[RMB]) // "3 ¥

r1 := [...]int{99: -1}
fmt.Println(r1)


a1 := [2]int{1, 2}
b1 := [...]int{1, 2}
c1 := [2]int{1, 3}
fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"
//d1 := [3]int{1, 2}
//fmt.Println(a1 == d1) // compile error: cannot compare [2]int == [3]int

var headHeaderKey = []byte("LastHeader")
fmt.Println(headHeaderKey)

}



