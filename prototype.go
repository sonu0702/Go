package main

import (
	"fmt"
)

//User obj
type User struct {
	email string
	phone string
}

//NewUser some
func NewUser(email string, phone string) User {
	return User{
		email: email,
		phone: phone,
	}
}

// WithEmail creates a copy of User with the provided email
func (u User) WithEmail(email string) User {
	u.email = email
	return u
}

// WithPhone creates a copy of User with the provided phone
func (u User) WithPhone(phone string) User {
	u.phone = phone
	return u
}

func main() {
	u := NewUser("email@gmail.com", "9876543")
	uu := u.WithEmail("email@hotmail.com")
	fmt.Println("Original User Object ", u)
	fmt.Println("User object copied during runtime ", uu)
}
