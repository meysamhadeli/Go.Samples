package variadic

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
