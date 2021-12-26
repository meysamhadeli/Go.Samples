package conditionals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	userAge, err := getUserAge()

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	if userAge >= 18 && userAge < 50 {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Sorry, Your age not enough")
	}
}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\r\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	return int(userAge), err
}
