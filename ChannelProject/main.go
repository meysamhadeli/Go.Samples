package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {

	c := make(chan int)

	go generateValue(c)
	go generateValue(c)

	x := <-c
	y := <-c

	sum := x + y
	fmt.Println(sum)
}

func generateValue(c chan int) int {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := randN.Intn(10)
	c <- result
	return result
}

//func main() {
//	greet()
//	stroreData("This is some dummy data!", "dummy-data.txt")
//
//	channel := make(chan int)
//
//	go storeMoreData(50000, "50000_1.txt", channel)
//	go storeMoreData(50000, "50000_2.txt", channel)
//
//	<-channel
//	<-channel
//}
//
//func greet() {
//	println("Hi there!")
//}
//
//func stroreData(storableText string, filename string) {
//	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//
//	if err != nil {
//		fmt.Println("Creating file failed!")
//	}
//
//	defer file.Close()
//
//	_, err = file.WriteString(storableText)
//
//	if err != nil {
//		fmt.Println("Writing to file failed!")
//	}
//}
//
//func storeMoreData(lines int, fileName string, c chan int) {
//	for i := 0; i < lines; i++ {
//		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
//		stroreData(text, fileName)
//	}
//	fmt.Printf("-- Done storing %v Lines of text --\n", lines)
//	c <- 1
//}
