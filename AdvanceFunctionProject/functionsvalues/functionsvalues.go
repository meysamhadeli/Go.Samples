package functionsvalues

import "fmt"

type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformFunction(&numbers)
	transformerFn2 := getTransformFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	transformedMoreNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(transformedMoreNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
