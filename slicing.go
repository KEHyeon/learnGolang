package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]
	fmt.Println("arrya :", array)
	fmt.Println("slice :", slice, len(slice), cap(slice))
	slice[0] = 100
	fmt.Println("arrya :", array)
	fmt.Println("slice :", slice, len(slice), cap(slice))
}
