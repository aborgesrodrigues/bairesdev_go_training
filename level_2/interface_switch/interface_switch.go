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

func GetName(n Namer) string {
	switch n.(type) {
	case *User:
		return "User " + n.Name()
	case *Customer:
		return "Customer " + n.Name()
	}
	return ""
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	c := &Customer{42, "Francesc"}

	fmt.Println(GetName(u))
	fmt.Println(GetName(c))
}
