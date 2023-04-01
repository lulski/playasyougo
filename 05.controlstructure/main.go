package main

import (
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	fmt.Println(quote.Go())
}
