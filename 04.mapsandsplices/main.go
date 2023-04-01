package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	//map
	myMap := make(map[string]string)

	myMap["dog"] = "samson"
	myMap["other-dog"] = "delilah"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMapInt := make(map[string]int)
	myMapInt["first"] = 1
	myMapInt["second"] = 2

	userMap := make(map[string]User)
	userMap["user1"] = User{
		FirstName: "trevor",
		LastName:  "Belmont",
	}

	log.Println(userMap["user1"])

	//var myNewVar float32

	//	myNewVar = 11.1

	//userMap

	//slice

	var mySlice []string
	mySlice = append(mySlice, "Nadya")

	mySlice = append(mySlice, "Bambi")
	mySlice = append(mySlice, "Rozi")

	log.Println(mySlice)

	log.Println(mySlice)

}
