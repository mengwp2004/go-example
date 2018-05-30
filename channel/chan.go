package main


import (
"fmt"
)


func main() {
	c := make(chan int, 2) //创建带有缓冲的chanel，缓冲大小是2


	//这样调用函数，那么f1和f2就是并发执行了
	go f1(c) //将参数c传递给f1()
	go f2(c) //将参数c传递给f2()


	c1 := <-c
	c2 := <-c //main函数只有从c中接收到俩个值，才会退出main()，否则main()中会阻塞这那直到c中有数据可以接收
	fmt.Printf("c1:%d   c2:%d", c1, c2)
}


func f1(c chan int) { // chan int 表示参数的类型是存储int类型的chanel
c <- 1 //向这个chanel中传入1，之后main()中就会接受到1
}


func f2(c chan int) { // chan int 表示参数的类型是存储int类型的chanel
c <- 2 //向这个chanel中传入2，之后main()中就会接收到2
}
