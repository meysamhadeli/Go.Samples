package main

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int
type personData map[string]person

const (
	PlayerChoiceAttack = iota + 1
	PlayerChoiceHeal
	PlayerChoiceSpecialAttack
)

func main() {

	//hobbies := []string{"Sport", "Reading"}

	//make
	hobbies := make([]string, 2, 10)
	hobbies[0] = "Sport"
	hobbies[1] = "Reading"

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)

	//new

	moreHobbies := new([]string)

	fmt.Println(*moreHobbies)

	*moreHobbies = append(*moreHobbies, "Sports")
	fmt.Println(*moreHobbies)

	//custom type

	var people personData = personData{
		"p1": {"Meysam", 33},
	}

	fmt.Println(people)

	var dummyNumber customNumber = 5
	poweredNumber := dummyNumber.pow(3)

	fmt.Println(poweredNumber)

	// special const
	fmt.Println(PlayerChoiceHeal)
	fmt.Println(PlayerChoiceSpecialAttack)
}

func (number customNumber) pow(power int) customNumber {
	result := number

	for i := 1; i < power; i++ {
		result = result * number
	}
	return result
}
