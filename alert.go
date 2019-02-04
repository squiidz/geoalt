package geoalt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/geo/s2"

	"github.com/dgraph-io/badger"
)

type Alert struct {
	ID        uint32
	CellID    s2.CellID
	Lat       float64
	Lng       float64
	UserID    uint32
	Message   string
	Timestamp string
}

func (a *Alert) Key(attr string) []byte {
	// alert:$alert_id:$user_id:$attribute_name = $value
	return []byte(fmt.Sprintf("alert:%d:%d:%d:%s", a.CellID, a.ID, a.UserID, attr))
}

func (a *Alert) SetAttr(attr string, value string) {
	switch attr {
	case "message":
		a.Message = value
	case "timestamp":
		a.Timestamp = value
	case "latitude":
		v := []byte(value)
		a.Lat = float64fromBytes(v)
	case "longitude":
		v := []byte(value)
		a.Lng = float64fromBytes(v)
	}
}

func (db *DB) GetAlert(cellID uint64, id uint32) (Alert, error) {
	var alert Alert

	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := []byte(fmt.Sprintf("alert:%d:%d", cellID, id))
	itr.Seek([]byte(pre))
	for itr.ValidForPrefix(pre) {
		keySplit := strings.Split(string(itr.Item().Key()), ":")
		attr := keySplit[len(keySplit)-1]
		itr.Item().Value(func(val []byte) error {
			alert.SetAttr(attr, string(val))
			return nil
		})
		uid, err := strconv.Atoi(keySplit[3])
		if err != nil {
			continue
		}
		alert.UserID = uint32(uid)
		itr.Next()
	}
	alert.ID = id
	alert.CellID = s2.CellID(cellID)
	return alert, nil
}

func (db *DB) GetUserAlertIDs(cellID uint64, id uint32) []int {
	var alertIDs []int

	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := []byte(fmt.Sprintf("alert:%d", cellID))
	itr.Seek([]byte(pre))
	for itr.ValidForPrefix(pre) {
		keySplit := strings.Split(string(itr.Item().Key()), ":")
		userID, err := strconv.Atoi(keySplit[3])
		if err != nil {
			continue
		}
		if uint32(userID) == id {
			aID, err := strconv.Atoi(keySplit[2])
			if err != nil {
				continue
			}
			if !contains(alertIDs, aID) {
				alertIDs = append(alertIDs, aID)
			}
		}
		itr.Next()
	}
	return alertIDs
}

func (db *DB) InsertAlert(a *Alert) error {
	txn := db.NewTransaction(true)
	a.ID = db.AlertSize()
	txn.Set(a.Key("message"), []byte(a.Message))
	txn.Set(a.Key("timestamp"), []byte(a.Timestamp))
	txn.Set(a.Key("latitude"), float64ToBytes(a.Lat))
	txn.Set(a.Key("longitude"), float64ToBytes(a.Lng))
	return txn.Commit()
}

func (db *DB) AlertSize() uint32 {
	var count uint32
	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	itr.Seek([]byte("alert"))
	for itr.ValidForPrefix([]byte("alert")) {
		count++
		itr.Next()
	}
	return count
}
