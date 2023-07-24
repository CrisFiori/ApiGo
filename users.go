package main

type userType struct {
	Id      string
	Name    string
	IsAdmin bool
}

var users = []userType{
	{Id: "1", Name: "Iron Man", IsAdmin: true},
	{Id: "2", Name: "Thor", IsAdmin: true},
	{Id: "3", Name: "Spiderman", IsAdmin: true},
}
