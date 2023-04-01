package main

import (
	"log"
	"time"
)

var s = "seven"

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// this function is called receiver function,
func (m *User) printFirstName() string {
	return m.FirstName
}

func main() {
	user1 := User{
		FirstName:   "Trevor",
		LastName:    "Belmont",
		PhoneNumber: "0474 123 258",
	}

	var user2 User
	user2.FirstName = "Alucard"
	user2.LastName = "Dracula"

	log.Println(user1.printFirstName(), user1.LastName, "Birthdate", user1.BirthDate)
	log.Println(user2.printFirstName(), user2.LastName, "Birthdate", user2.BirthDate)

}
