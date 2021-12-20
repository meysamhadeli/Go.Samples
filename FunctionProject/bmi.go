package main

import (
	"github.com/meysamhadeli/Go.Samples/FunctionProject/info"
	"github.com/meysamhadeli/Go.Samples/FunctionProject/reader"
)

func main() {
	info.PrintWelcome()

	weight, height := reader.GetUserMetrics()
	bmi := calculateBMI(weight, height)
	PrintBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
