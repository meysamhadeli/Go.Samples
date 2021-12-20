package main

import (
	"fmt"
	"github.com/meysamhadeli/Go.Samples/VariableProject/info"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Seprator)
	fmt.Print("Please enter your weight (kg):")

	weightInput, _ := info.Reader.ReadString('\n')
	fmt.Print("Please enter your height (m):")
	heightInput, _ := info.Reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Print("Your BMI: ", bmi)

}
