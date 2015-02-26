package main

import "fmt"
import "strconv"

type Ishan struct {
	nos     int
	welcome string
}

type MyPrinterType func(string)

func Greet(person Ishan, myWassa MyPrinterType) {
	myGreetingMas := CreateMessage(person.nos, person.welcome, "howdy")
	myWassa(myGreetingMas)
}

func CreateMessage(age int, welcomeMas ...string) (myGreeting string) {
	switch "Blah" {
	case "Tim":
		myGreeting = ("Deffered Message =>")
	case "Blah":
		myGreeting = ("Added Message =>")
	}
	if welcomeMas[1] != "" {
		myGreeting = myGreeting + welcomeMas[1] + " of age " + strconv.Itoa(age)
	} else {
		myGreeting = myGreeting + welcomeMas[0] + " of age " + strconv.Itoa(age)
	}
	return
}

func myPrintFunction(custom string) MyPrinterType {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {
	var t = Ishan{21, "Blah"}
	Greet(t, myPrintFunction("!!!"))
}
