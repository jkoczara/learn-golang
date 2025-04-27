package main

import (
	"fmt"

	"example.com/structs/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@test.com", "abcd")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	var name str = "Max"
	name.log()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
