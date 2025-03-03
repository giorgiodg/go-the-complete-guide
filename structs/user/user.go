package user

import (
	"time"
	"fmt"
	"errors"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}


type Admin struct {
	email string
	password string
	User User
}

// method of this struct
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// method of this struct
func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}


func NewAdmin(email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName :"Admin",
			lastName :"Admin",
			birthdate :"---",
			createdAt : time.Now(),
		},
	}
}

// this is a best practice/convention
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil

}