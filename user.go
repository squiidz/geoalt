package geoalt

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dgraph-io/badger"
)

type UserStore struct {
	*badger.DB
}

type User struct {
	ID        uint32
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

func (u *User) KeyEmail() []byte {
	// user:$email= $id
	return []byte(fmt.Sprintf("user:%s", u.Email))
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

func (db *UserStore) GetUserIDs(attr, value string) []int {
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

func (db *UserStore) GetUser(id uint32) (*User, error) {
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
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("User not found")
	}
	user.ID = uint32(id)
	return &user, nil
}

func (db *UserStore) GetUserByEmail(email string) (*User, error) {
	var id uint32
	txn := db.NewTransaction(false)
	key := []byte(fmt.Sprintf("user:%s", email))

	itm, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	itm.Value(func(val []byte) error {
		id = uint32FromBytes(val)
		return nil
	})
	return db.GetUser(id)
}

func (db *UserStore) Insert(u *User) error {
	if db.exist(u) {
		return errors.New("User already exist")
	}
	txn := db.NewTransaction(true)
	u.ID = db.NextID()
	txn.Set(u.Key("email"), []byte(u.Email))
	txn.Set(u.Key("password"), []byte(u.Password))
	txn.Set(u.Key("first_name"), []byte(u.FirstName))
	txn.Set(u.Key("last_name"), []byte(u.LastName))
	txn.Set(u.Key("address"), []byte(u.Address))
	txn.Set(u.KeyEmail(), uint32ToBytes(u.ID))
	return txn.Commit()
}

func (db *UserStore) NextID() uint32 {
	txn := db.NewTransaction(true)
	defer txn.Commit()
	lenKey := []byte("user:len")
	ulen := db.Size() + 1
	if ulen == 0 {
		txn.Set(lenKey, uint32ToBytes(0))
	}
	err := txn.Set(lenKey, uint32ToBytes(ulen))
	if err != nil {
		log.Fatal(err)
	}
	return ulen
}

func (db *UserStore) Size() uint32 {
	var size uint32
	txn := db.NewTransaction(false)
	itm, err := txn.Get([]byte("user:len"))
	if err != nil {
		return 0
	}
	err = itm.Value(func(val []byte) error {
		size = uint32FromBytes(val)
		return nil
	})
	return size
}

func (db *UserStore) exist(u *User) bool {
	txn := db.NewTransaction(false)
	itm, err := txn.Get(u.KeyEmail())
	if err != nil || itm == nil {
		return false
	}
	return true
}
