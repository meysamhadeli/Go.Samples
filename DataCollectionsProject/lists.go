package main

import "fmt"

func main() {
	var productNames = [5]string{"book1", "book2", "book3", "book4"}
	productNames[4] = "last book"
	fmt.Println(productNames)

	productNamesSlice := productNames[1:3]
	fmt.Println(productNamesSlice)
	fmt.Println(len(productNamesSlice), cap(productNamesSlice))
}
