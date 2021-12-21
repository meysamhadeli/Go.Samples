package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func main() {

	createdProduct := getProduct()

	createdProduct.printData()
	createdProduct.store()
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

func (prod *Product) printData() {
	fmt.Println(*prod)
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("-----------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Id: ")
	titleInput := readUserInput(reader, "Title: ")
	descriptionInput := readUserInput(reader, "Description: ")
	priceInput := readUserInput(reader, "Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)
	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	return userInput
}

func (prod *Product) store() {
	file, _ := os.Create(prod.id + ".txt")
	content := fmt.Sprintf("Id: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		prod.id, prod.title, prod.description, prod.price)

	file.WriteString(content)
	file.Close()
}
