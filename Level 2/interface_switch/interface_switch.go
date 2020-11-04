package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func printName(n Namer) {
	switch n.(type) {
	case *User:
		fmt.Println("User", n.Name())
	case *Customer:
		fmt.Println("Customer", n.Name())
	}
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	c := &Customer{42, "Francesc"}

	printName(u)
	printName(c)
}
