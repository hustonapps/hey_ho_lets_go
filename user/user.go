package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User // embedded struct
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin { // returns a copy of Admin
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "-/-/-",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) { // return a pointer to User
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	return &User{
		firstName, 
		lastName, 
		birthdate, 
		time.Now(),
	}, nil
}