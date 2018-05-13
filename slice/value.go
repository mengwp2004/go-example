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

    var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var mySlice []int = myArray[:5]  // Go语言支持用myArray[first:last]格式选择数组元素
    fmt.Println(mySlice);
    mySlice = append(mySlice,1,2,3,4)
    fmt.Println(mySlice)

    fmt.Println("slice:")
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{9, 8, 7}

    copy(slice2, slice1)  // 复制slice1的前三个元素到slice2中
    copy(slice1, slice2)  // 复制slice2的3个元素到slice2的前3个位置
    fmt.Println(slice1)
    fmt.Println(slice2)
}
