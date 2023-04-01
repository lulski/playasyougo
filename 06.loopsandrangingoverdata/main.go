package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "horse", "elephant", "tiger"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		age       int
	}
	var users []User
	users = append(users, User{"John1", "Smith1", "johnsmit1@abc.com", 30})
	users = append(users, User{"John2", "Smith2", "johnsmit2@abc.com", 40})
	users = append(users, User{"John3", "Smith3", "johnsmit3@abc.com", 50})
	users = append(users, User{"John4", "Smith4", "johnsmit4@abc.com", 60})

	for i, l := range users {
		log.Println(i, l)
	}

}
