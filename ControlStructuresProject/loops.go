package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	selectedChoice, err := getUserChoice()
	if err != nil {
		fmt.Println("invalid choice, exiting!")
		return
	}

	if selectedChoice == "1" {
		calculateSumUpToNumber()

	} else if selectedChoice == "2" {
		calculateFactorial()

	} else if selectedChoice == "3" {
		calculateSumManually()

	} else {
		calculateListSum()
	}
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\r\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		fmt.Println("Invalid input")
		return 0, err
	}

	return int(chosenNumber), nil
}

func calculateSumUpToNumber() {
	fmt.Print("Please enter your number:  ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	fmt.Println(chosenNumber)
	sum := 0

	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}

	println("Result: %v", sum)
}

func calculateFactorial() {
	fmt.Print("Please enter your number:  ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	factorial := 1

	for i := 1; i <= chosenNumber; i++ {
		factorial = factorial * i
	}

	println("Result: %v", factorial)
}

func calculateSumManually() {
	sum := 0
	isEnteringNumbers := true

	println("Keep on entering numbers")

	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum = sum + chosenNumber

		if err != nil {
			isEnteringNumbers = false
		}
	}

	fmt.Println("Result: %v\n", sum)

}

func calculateListSum() {
	fmt.Println("Please enter a comma-separated list of numbers.")

	inputNumberList, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\r\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, ",")

	sum := 0

	for index, value := range inputNumbers {
		fmt.Println("Index: %v, Value: %v\n", index, value)
		number, err := strconv.ParseInt(value, 0, 64)

		if err != nil {
			continue
		}
		sum = sum + int(number)
	}

	fmt.Printf("Result: %v\n", sum)
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")

	fmt.Print("please make your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\r\n", "", -1)
	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {

		return userInput, nil
	} else {
		return "", errors.New("Invalid input!")
	}
}
