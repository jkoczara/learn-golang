package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email     string
	passsword string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birth date are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:     email,
		passsword: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		}
	}
}
