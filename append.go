package main
import "fmt"

func main(){

slice1 := make([]int, 5, 10)
slice2 := append(slice1, 5, 6, 7, 8, 9)
fmt.Println(slice2, len(slice2), cap(slice2))


slice2 = append(slice2, 10) 
fmt.Println(slice2, len(slice2), cap(slice2))



    d := []string{"Welcome", "for", "Hangzhou, ", "Have", "a", "good", "journey!"}
    insertSlice := []string{"It", "is", "a", "big", "city, "}
    insertSliceIndex := 3
    d = append(d[:insertSliceIndex], append(insertSlice, d[insertSliceIndex:]...)...)
    fmt.Println(d)

var runes []rune
for _, r := range "Hello, 世界" {
    runes = append(runes, r)
}
fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
fmt.Printf("%T\n",runes)

var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // append the slice x
fmt.Println(x)
// "[1 2 3 4 5 6 1 2 3 4 5 6]"

fmt.Println(sum()) // "0"
fmt.Println(sum(3)) // "3"
fmt.Println(sum(1, 2, 3, 4)) // "10"


values := []int{1, 2, 3, 4}
//fmt.Println(sum(values...)) // "10"
fmt.Println(sum(values...)) // "10"



}

func sum(vals...int) int {
total := 0
for _, val := range vals {
total += val
}
return total
}
