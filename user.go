package geoalt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgraph-io/badger"
)

type User struct {
	ID        int
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Address   string `db:"address"`
}

func (u *User) Key(attr string) []byte {
	// user:$user_id:$attribute_name = $value
	return []byte(fmt.Sprintf("user:%d:%s", u.ID, attr))
}

func (u *User) SetAttr(attr string, value string) {
	switch attr {
	case "email":
		u.Email = value
	case "password":
		u.Password = value
	case "first_name":
		u.FirstName = value
	case "last_name":
		u.LastName = value
	case "address":
		u.Address = value
	}
}

var fakeUsers = []User{
	{ID: 1, Email: "john@mail.com", Password: "a1b2c3d4", FirstName: "John", LastName: "Doe", Address: "New York"},
	{ID: 2, Email: "ben@mail.com", Password: "1234abcd", FirstName: "Benjamin", LastName: "Button", Address: "Chicago"},
	{ID: 3, Email: "mycroft@mail.com", Password: "abcdef", FirstName: "Mycroft", LastName: "Holmes", Address: "London"},
	{ID: 4, Email: "fee@mail.com", Password: "123456", FirstName: "Fee", LastName: "Courtemanche", Address: "Quebec"},
	{ID: 5, Email: "jo@mail.com", Password: "qwerty", FirstName: "Jonathan", LastName: "Chaput", Address: "Quebec"},
}

func (db *DB) GetUserIDs(attr, value string) []int {
	var userIDs []int

	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	itr.Seek([]byte("user"))
	for itr.Valid() {
		key := string(itr.Item().Key())
		keySplit := strings.Split(key, ":")
		uID, err := strconv.Atoi(keySplit[1])
		if err != nil {
			continue
		}
		if keySplit[len(keySplit)-1] == attr {
			itr.Item().Value(func(val []byte) error {
				if string(val) == value {
					userIDs = append(userIDs, uID)
				}
				return nil
			})
		}
		if keySplit[0] != "user" {
			break
		}
		itr.Next()
	}
	return userIDs
}

func (db *DB) GetUser(id int) (User, error) {
	var user User

	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := []byte(fmt.Sprintf("user:%d", id))
	itr.Seek([]byte(pre))
	for itr.ValidForPrefix(pre) {
		keySplit := strings.Split(string(itr.Item().Key()), ":")
		attr := keySplit[len(keySplit)-1]
		itr.Item().Value(func(val []byte) error {
			user.SetAttr(attr, string(val))
			return nil
		})
		itr.Next()
	}
	user.ID = id
	return user, nil
}

func (db *DB) InsertUser(u *User) error {
	txn := db.NewTransaction(true)
	txn.Set(u.Key("email"), []byte(u.Email))
	txn.Set(u.Key("password"), []byte(u.Password))
	txn.Set(u.Key("first_name"), []byte(u.FirstName))
	txn.Set(u.Key("last_name"), []byte(u.LastName))
	txn.Set(u.Key("address"), []byte(u.Address))
	return txn.Commit()
}
