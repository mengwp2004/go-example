package main

import(
"fmt"
)

func modify(arr [5]int){
   arr[0] = 10
   fmt.Println("In modify(), arr values:", arr)
}

func main(){
    arr := [5]int{1, 2, 3, 4, 5}
    modify(arr)
    fmt.Println("In main(), arr values:", arr)
}
