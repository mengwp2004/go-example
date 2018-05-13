package main

import "fmt"

func main() {
	x := 1
	p := &x
	// p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	v := 1
	incr(&v)
	// side effect: v is now 2fmt.Println(incr(&v)) // "3" (and v is 3)
	fmt.Println(v)

	q := new(int)
	// p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*q) // "0"
	*q = 2
	// 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
	var a = b + c   // a 第三个初始化, 为 3
	var b = f()     // b 第二个初始化, 为 2, 通过调用 f (依赖c)
	var c = 1       // c 第一个初始化, 为 1

	fmt.Println(a)

}
func f() int { return c + 1 }

func incr(p *int) int {
	*p++ // 非常重要:只是增加p指向的变量的值,并不改变p指针!!!
	return *p
}
