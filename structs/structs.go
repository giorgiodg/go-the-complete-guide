package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	// u is an instance of the User struct, which represents a user with specific attributes.
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		return
	}

	admin := user.NewAdmin("test@test.com", "test123")
	admin.User.OutputUserDetails()
	admin.User.ClearUsername()
	admin.User.OutputUserDetails()

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
