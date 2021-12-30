package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user-log.txt"}
	createLog(userLog)

	message := loggableString("[Info] user created!")
	createLog(message)

	outputValue(userLog)
	outputValue(message)

	firstNumber := 5
	secondNumber := 10.1
	numbers := []int{1, 2, 3, 4, 5}

	doubleFirstNumber := double(firstNumber)
	doubleSecondNumber := double(secondNumber)
	doubleThirdNumber := double(numbers)
	doubleFourthNumber := double("Test")

	fmt.Println(doubleFirstNumber)
	fmt.Println(doubleSecondNumber)
	fmt.Println(doubleThirdNumber)
	fmt.Println(doubleFourthNumber)
}

func createLog(data logger) {
	data.log()
}

func outputValue(value interface{}) {
	fmt.Println(value)
}

func double(value interface{}) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		newNumbers := append(val, val...)
		return newNumbers
	default:
		return ""
	}
}
