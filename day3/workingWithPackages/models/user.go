package models

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

func (u User) Display(){
	 fmt.Printf("User: %s (ID: %d) - %s\n", u.Name, u.Id, u.Email)
}

func NewUser(id int, name, email string) User {
    return User{
        Id:    id,
        Name:  name,
        Email: email,
    }
}