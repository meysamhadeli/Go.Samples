package reader

import (
	"bufio"
	"fmt"
	"github.com/meysamhadeli/Go.Samples/FunctionProject/info"
	"os"
	"strconv"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

func GetUserMetrics() (weight float64, height float64) {

	weight = GetUserInput(info.HeightPrompt)
	height = GetUserInput(info.WeightPrompt)
	return
}

func GetUserInput(promptText string) float64 {

	fmt.Print(promptText)
	userInput, _ := Reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	value, _ := strconv.ParseFloat(userInput, 64)

	return value
}
