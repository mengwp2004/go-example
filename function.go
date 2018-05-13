package main
import "fmt"
import "math"
func hypot(x, y float64) float64 {
   return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int {return x + y}
func sub(x, y int) (z int) { z = x - y; return}
func first(x int, _ int) int { return x }
func zero(int, int) int{ return 0 }
func (l int) zero1() int{ 
  fmt.Println(l) 
   return 0 
}


func main(){

   fmt.Println(hypot(3,4)) // "5"
fmt.Printf("%T\n", add) // "func(int, int) int"
fmt.Printf("%T\n", sub) // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero) // "func(int, int) int"
fmt.Printf("%T\n", zero1) // "func(int, int) int"



}
