package main

type User struct {
	ID    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		ID:    "1",
		Name:  "木兮",
		Phone: "1234567890",
	},
	"2": {
		ID:    "1",
		Name:  "小慕",
		Phone: "1234567890",
	},
}
