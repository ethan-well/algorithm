package main

import "fmt"

func someChage(slice *[]int) {
	fmt.Println(slice)
	fmt.Printf("addres2: %p\n", *slice)

	*slice = append(*slice, 1, 2, 3)
	fmt.Printf("addres3: %p\n", *slice)

	*slice = append(*slice, 4, 5, 6)
	fmt.Printf("addres4: %p\n", *slice)
	fmt.Println(slice)
}

func main() {
	slice := make([]int, 0, 10)
	fmt.Printf("addres1: %p\n", slice)

	someChage(&slice)

	fmt.Printf("addres5: %p\n", slice)
	fmt.Println(slice)
}
