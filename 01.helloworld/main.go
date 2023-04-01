package main

import "fmt"

func main() {
	fmt.Println("hello world!")

	var whatToSay string
	var i int

	whatToSay = "goodbye cruel world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
