package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main Started")
	fmt.Println("For Loop")
	for i := 0; i < 30; i = i + 3 {
		fmt.Println(i)
	}

	fmt.Println("Continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
	}

	mySliceEntry := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 110}

	for j, slicePointer := range mySliceEntry {
		fmt.Println(j, " - ", slicePointer)
	}
	fmt.Print("[12:]")
	fmt.Println(mySliceEntry[1:])

	//string slicing
	myStringSlice := "Slice cutting Here !!!"
	fmt.Println(myStringSlice)
	fmt.Println(myStringSlice[6:])

}
