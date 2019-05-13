package model

type Account struct {
	Base
	Name string
	UUID string
	Book []Book
}

type Admin struct {
	Account
	Category    []Category
	Author      []Author
	Description []Description
}
