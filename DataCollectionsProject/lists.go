package main

import "fmt"

//func main() {
//	var productNames = [5]string{"book1", "book2", "book3", "book4"}
//	productNames[4] = "last book"
//	fmt.Println(productNames)

//	productNamesSlice := productNames[1:3]
//	fmt.Println(productNamesSlice)
//	fmt.Println(len(productNamesSlice), cap(productNamesSlice))
//}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 8.12)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

}
