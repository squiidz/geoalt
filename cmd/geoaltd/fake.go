package main

import "github.com/squiidz/geoalt"

var fakeUsers = []geoalt.User{
	{ID: 1, Email: "john@mail.com", Password: "a1b2c3d4", FirstName: "John", LastName: "Doe", Address: "New York"},
	{ID: 2, Email: "ben@mail.com", Password: "1234abcd", FirstName: "Benjamin", LastName: "Button", Address: "Chicago"},
	{ID: 3, Email: "mycroft@mail.com", Password: "abcdef", FirstName: "Mycroft", LastName: "Holmes", Address: "London"},
	{ID: 4, Email: "fee@mail.com", Password: "123456", FirstName: "Fee", LastName: "Courtemanche", Address: "Quebec"},
	{ID: 5, Email: "jo@mail.com", Password: "qwerty", FirstName: "Jonathan", LastName: "Chaput", Address: "Quebec"},
}

func PopulateDB(db *geoalt.DB) error {
	for _, u := range fakeUsers {
		err := db.UserStore.Insert(&u)
		if err != nil {
			return err
		}
	}
	return nil
}
