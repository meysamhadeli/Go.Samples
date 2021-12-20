package info

import (
	"fmt"
)

const MainTitle = "BMI Calculatore"
const Seprator = "-------------------"
const WeightPrompt = "Please enter your weight (kg):"
const HeightPrompt = "Please enter your height (m):"

func PrintWelcome() {
	fmt.Println(MainTitle)
	fmt.Println(Seprator)
}
