package main

import (
	"fmt"

	"github.com/hustonapps/m/v2/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return ;
	}

	admin := user.NewAdmin("test@test.com", "password")
	admin.OutputUserDetails()
	admin.ClearUsername()
	admin.User.OutputUserDetails() // another way of calling

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
