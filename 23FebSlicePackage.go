package main

import "fmt"
import "Web_Programming_CSUF/Web_Programming_CSUF/src/github.com/goestoeleven/math"

type myMas string

// TYPES CAN CONTAIN METHODS
func (s myMas) messagePrint() string {
	return "METHOD: Printing Message"
}

func main() {

	numberSlice := []int{5, 10, 15}

	sumOfNumbers := math.Sum(numberSlice)

	fmt.Println(sumOfNumbers)

	mySlice := []int{1, 5, 15, 20, 25, 30}
	myOtherSlice := []int{50, 70, 90}

	mySlice = append(mySlice, myOtherSlice...)

	for i, currentEntry := range mySlice {
		fmt.Println(i, " - ", currentEntry)
	}

	fmt.Println("")

	mySlice = append(mySlice[:4], mySlice[5:]...)

	for i, currentEntry := range mySlice {
		fmt.Println(i, " - ", currentEntry)
	}

	var mas myMas = "It Works"
	fmt.Println(mas.messagePrint())

}
