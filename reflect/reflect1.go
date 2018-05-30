package main

import (
   "fmt"
   "reflect"
)

type Foo struct {
  X string
  Y int
}

func (f Foo) do() {
  fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)

}

func main() {
	var s string = "abc"
	fmt.Println(reflect.TypeOf(s).String()) //string
	fmt.Println(reflect.TypeOf(s).Name())   //string
        _ = "breakpoint"
	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ.String()) //main.Foo
	fmt.Println(typ.Name())     //Foo ，返回结构体的名字

	for i := 0; i < typ.NumField(); i++ {
	field := typ.Field(i)
	fmt.Printf("%s type is :%s\n", field.Name, field.Type)
	}

	//x type is :string
	//y type is :int

	field2, _ := typ.FieldByName("X") //等价于typ.Field(0)，返回的也是StructField对象
	fmt.Println(field2.Name)          // x

	fmt.Println(typ.NumMethod()) //1， Foo 方法的个数
        if (typ.NumMethod() > 0 ) {
		m := typ.Method(0)
		fmt.Println(m.Name) //do
		fmt.Println(m.Type) //func(main.Foo)
		fmt.Println(m.Func)
        }

	fmt.Println(typ.Kind()) //struct

	var f2 = &Foo{}
	typ2 := reflect.TypeOf(f2)
	fmt.Println(typ2)        //*main.Foo
	fmt.Println(typ2.Kind()) //ptr

	var i int = 123
	var f1 = Foo{"abc1", 1235}
	fmt.Println(reflect.ValueOf(i)) //<int Value>
	fmt.Println(reflect.ValueOf(f1)) //<main.Foo Value>
	fmt.Println(reflect.ValueOf(s)) //abc



val := reflect.ValueOf(f1)
fmt.Println(val.FieldByName("Y")) //<int Value>  interface.Value对象

fmt.Println(typ.FieldByName("Y")) //{  <nil>  0 [] false} false StructField对象

started := []reflect.Type{}

typ1 :=reflect.TypeOf(started)
        fmt.Println(typ1.String()) //main.Foo
        fmt.Println(typ1.Name()) 
}

