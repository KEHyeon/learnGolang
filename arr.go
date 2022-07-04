package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300
	for _, v := range nums {
		fmt.Println(v)
	}
}
