package main

import (
	"errors"
	"fmt"
	"strings"
)

type user struct {
	name  string
	age   uint8
	email string
}

// init map globally for future use in findUser() function
var usersMap map[string]user = make(map[string]user)

func main() {
	users := []user{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}

	// fill our map for the future users search
	for _, user := range users {
		key := strings.ToLower(user.name)
		usersMap[key] = user
	}

	// find user "abobus"
	err, ourUser := findUser("abobus")

	// show error or user info depending on result
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Name: %s, Age: %d, Email: %s", ourUser.name, ourUser.age, ourUser.email)
	}
}

func findUser(name string) (error, user) {
	name = strings.ToLower(name)
	value, found := usersMap[name]

	if !found {
		error := errors.New("User not found")
		return error, value
	}

	return nil, value
}
