package main

import (
	"fmt"
	"os"
)

func main() {
	greet()
	stroreData("This is some dummy data!", "dummy-data.txt")
}

func greet() {
	println("Hi there!")
}

func stroreData(storableText string, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Creating file failed!")
	}

	defer file.Close()

	_, err = file.WriteString(storableText)

	if err != nil {
		fmt.Println("Writing to file failed!")
	}

}
