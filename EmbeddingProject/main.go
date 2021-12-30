package main

import (
	"bufio"
	"fmt"
	"os"
)

type Storer interface {
	Store(fileName string)
}

type Prompter interface {
	PromptForInput()
}

type PromptStorer interface {
	Storer
	Prompter
}

type userInputData struct {
	input string
}

type user struct {
	firstName string
	lastName  string
	*userInputData
}

func main() {
	data := newUserInputData()
	meysam := newUser("meysam", "h")
	//data.PromptForInput()
	//data.Store("user1 .txt")

	meysam.PromptForInput()
	meysam.Store("meysam .txt")

	handleUserInput(data)
}

func newUser(first string, last string) *user {
	return &user{
		firstName:     first,
		lastName:      last,
		userInputData: &userInputData{},
	}
}

func newUserInputData() *userInputData {
	return &userInputData{}
}

func (user *userInputData) PromptForInput() {
	fmt.Print("Your input please: ")

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Fetching user input failed!")
		return
	}

	user.input = userInput
}

func (user *userInputData) Store(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Creating file failed!")
		return
	}

	defer file.Close()

	file.WriteString(user.input)
}

func handleUserInput(container PromptStorer) {
	fmt.Println("Ready to store your data")
	container.PromptForInput()
	container.Store("user2.txt")
}
