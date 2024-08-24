package models

type User struct {
	Email    string
	Password []byte
}

var AllUsers []User
