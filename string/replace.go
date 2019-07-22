package main

import (

     "encoding/hex"
     "fmt"
     "strings"
)

func main(){
 str := "00  11"
 str1 :=strings.Replace(str, " ", "", -1)
 b, _ := hex.DecodeString(strings.Replace(str, " ", "", -1))
 fmt.Println(str1)
 fmt.Printf("%v",b)
}
