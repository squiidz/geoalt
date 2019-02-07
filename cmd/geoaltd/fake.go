package main

import "github.com/squiidz/geoalt"

var fakeUsers = []geoalt.User{
	{ID: 1, Email: "john@mail.com", Password: "a1b2c3d4", FirstName: "John", LastName: "Doe", Address: "New York"},
	{ID: 2, Email: "ben@mail.com", Password: "1234abcd", FirstName: "Benjamin", LastName: "Button", Address: "Chicago"},
	{ID: 3, Email: "mycroft@mail.com", Password: "abcdef", FirstName: "Mycroft", LastName: "Holmes", Address: "London"},
	{ID: 4, Email: "fee@mail.com", Password: "123456", FirstName: "Fee", LastName: "Courtemanche", Address: "Quebec"},
	{ID: 5, Email: "jo@mail.com", Password: "qwerty", FirstName: "Jonathan", LastName: "Chaput", Address: "Quebec"},
}

var fakeAlerts = []geoalt.Alert{
	{ID: 1, CellID: 12345, UserID: 1, Message: "50000", Timestamp: "01/12/2017"},
	{ID: 2, CellID: 23456, UserID: 1, Message: "20000", Timestamp: "01/13/2017"},
	{ID: 3, CellID: 34567, UserID: 2, Message: "75000", Timestamp: "01/14/2017"},
	{ID: 4, CellID: 45678, UserID: 3, Message: "40000", Timestamp: "01/15/2017"},
	{ID: 5, CellID: 56789, UserID: 3, Message: "20000", Timestamp: "01/17/2017"},
	{ID: 6, CellID: 67890, UserID: 3, Message: "25000", Timestamp: "01/18/2017"},
	{ID: 7, CellID: 78901, UserID: 4, Message: "10000", Timestamp: "01/21/2017"},
	{ID: 8, CellID: 89012, UserID: 5, Message: "15000", Timestamp: "01/22/2017"},
	{ID: 9, CellID: 90123, UserID: 5, Message: "50000", Timestamp: "01/23/2017"},
}

func PopulateDB(db *geoalt.DB) error {
	for _, u := range fakeUsers {
		err := db.UserStore.Insert(&u)
		if err != nil {
			return err
		}
	}
	for _, a := range fakeAlerts {
		err := db.AlertStore.Insert(&a)
		if err != nil {
			return err
		}
	}
	return nil
}
