package main

import "fmt"

type Person struct {
	name string
	age int
	addr string
}

type Employee struct {
	Person          //匿名字段
	salary int
	int             //用内置类型作为匿名字段
	addr string     //类似于重载
}

func main() {
	/*
	var em1 Employee = Employee{}
	em1.Person = Person{"rain", 23, "qingyangqu"}
	em1.salary = 5000
	em1.int = 100 //使用时注意其意义，此处无
	em1.addr = "gaoxingqu"
	*/
	//em1 := Employee{Person{"rain", 23, "qingyangqu"}, 5000, 100, "gaoxingqu"}
	//初始化方式不一样，但是结果一样
	em1 := Employee{Person:Person{"rain", 23, "qingyangqu"}, salary:5000, int:100, addr:"gaoxingqu"}
	fmt.Println(em1)
	/******************************************************/
	/*print result                                        */
	/*{{rain 23 qingyangqu} 5000 100 gaoxingqu}           */
	/******************************************************/
	
	fmt.Println("live addr(em1.addr) = ", em1.addr)
	fmt.Println("work addr(em1.Person.addr) = ", em1.Person.addr)
	em1.int = 200  //修改匿名字段的值
	/******************************************************/
	/*print result                                        */
	/*live addr(em1.addr) =  gaoxingqu                    */
	/*work addr(em1.Person.addr) =  qingyangqu            */
	/******************************************************/
}
