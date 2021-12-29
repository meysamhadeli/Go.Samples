package recursion

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
