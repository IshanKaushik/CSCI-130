package main

import "fmt"

const (
	Study = "MS"
	City  = "Fresno"
)

type Ishan struct {
	name string
	age  int
}

func main() {

	var c = Ishan{"Ron", 22}

	MyFunc(c)

}

func MyFunc(ishan Ishan) {
	fmt.Println(ishan.name)
	fmt.Println(ishan.age)

	fmt.Println(Study)
	fmt.Println(City)

	fmt.Println(ReturnMessage())
	name, _ := ReturnMessage()
	fmt.Println(name)
}

func ReturnMessage() (string, int) {
	return "John", 12
}
