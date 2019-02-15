package geoalt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
	h3 "github.com/uber/h3-go"
)

type AlertStore struct {
	*badger.DB
}

type Alert struct {
	ID      uint32
	UserID  uint32
	Cell    Cell
	Coord   Coord
	Message string
	// Created at
	Timestamp int64
	// Last read
	ReadAt int64
	// Delay until avaible again
	Delay     int64
	Ephemeral bool
}

type Cell struct {
	// Smallest cell resolution (15)
	Base h3.H3Index
	// Cell id used for indexing
	Index h3.H3Index
	// Cell id with the resolution
	Real       h3.H3Index
	Resolution uint32
}

type Coord struct {
	Lat float64
	Lng float64
}

func (c *Coord) CellID(res int) h3.H3Index {
	return h3.FromGeo(h3.GeoCoord{Latitude: c.Lat, Longitude: c.Lng}, res)
}

func (a *Alert) Key(attr string) []byte {
	// alert:$alert_id:$user_id:$attribute_name = $value
	return []byte(fmt.Sprintf("alert:%d:%d:%d:%s", a.UserID, a.Cell.Index, a.ID, attr))
}

func (a *Alert) SetAttr(attr string, value []byte) {
	switch attr {
	case "message":
		a.Message = string(value)
	case "timestamp":
		a.Timestamp = int64FromBytes(value)
	case "read_at":
		a.ReadAt = int64FromBytes(value)
	case "delay":
		a.Delay = int64FromBytes(value)
	case "latitude":
		a.Coord.Lat = float64fromBytes(value)
	case "longitude":
		a.Coord.Lng = float64fromBytes(value)
	case "ephemeral":
		a.Ephemeral = boolFromBytes(value)
	case "base_cell":
		a.Cell.Base = h3.H3Index(uint64FromBytes(value))
	case "index_cell":
		a.Cell.Index = h3.H3Index(uint64FromBytes(value))
	case "real_cell":
		a.Cell.Real = h3.H3Index(uint64FromBytes(value))
	case "resolution":
		a.Cell.Resolution = uint32FromBytes(value)
	}
}

func (a *Alert) Borders() []*Coord {
	var coords []*Coord
	boundaries := h3.ToGeoBoundary(a.Cell.Real)
	for _, b := range boundaries {
		coords = append(coords, &Coord{
			Lat: b.Latitude,
			Lng: b.Longitude,
		})
	}
	return coords
}

func (a *Alert) ValidDelay() bool {
	return (a.ReadAt + a.Delay) < time.Now().Unix()
}

func (db *AlertStore) GetAlert(cellID uint64, userID uint32, id uint32) (*Alert, error) {
	var alert Alert
	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := []byte(fmt.Sprintf("alert:%d:%d:%d", userID, cellID, id))

	itr.Seek([]byte(pre))
	for itr.ValidForPrefix(pre) {
		keySplit := strings.Split(string(itr.Item().Key()), ":")
		attr := keySplit[len(keySplit)-1]
		itr.Item().Value(func(val []byte) error {
			alert.SetAttr(attr, val)
			return nil
		})
		itr.Next()
	}
	if !alert.ValidDelay() {
		return nil, errors.New("Delay invalid")
	}
	alert.ID = id
	alert.UserID = userID
	alert.Cell.Index = h3.H3Index(cellID)

	db.PutAttr(&alert, "read_at", int64ToBytes(time.Now().Unix()))
	if alert.Ephemeral {
		db.Delete(&alert)
	}
	return &alert, nil
}

func (db *AlertStore) GetUserAlertIDs(cellID uint64, userID uint32) []int {
	var alertIDs []int

	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := basePrefix(userID, cellID)
	itr.Seek([]byte(pre))
	for itr.ValidForPrefix(pre) {
		keySplit := strings.Split(string(itr.Item().Key()), ":")
		aID, err := strconv.Atoi(keySplit[3])
		if err != nil {
			continue
		}
		if !contains(alertIDs, aID) {
			alertIDs = append(alertIDs, aID)
		}
		itr.Next()
	}
	return alertIDs
}

func (db *AlertStore) Insert(a *Alert) error {
	if db.Exist(a) {
		return errors.New("Alert already exist")
	}
	var err error
	txn := db.NewTransaction(true)

	a.ID = db.Size(a.UserID, uint64(a.Cell.Index)) + 1
	a.Cell.Base = h3.FromGeo(h3.GeoCoord{Latitude: a.Coord.Lat, Longitude: a.Coord.Lng}, 15)

	if err = txn.Set(a.Key("message"), []byte(a.Message)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("timestamp"), int64ToBytes(a.Timestamp)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("read_at"), int64ToBytes(a.Timestamp)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("delay"), int64ToBytes(a.Delay)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("latitude"), float64ToBytes(a.Coord.Lat)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("longitude"), float64ToBytes(a.Coord.Lng)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("ephemeral"), boolToBytes(a.Ephemeral)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("base_cell"), uint64ToBytes(uint64(a.Cell.Base))); err != nil {
		return err
	}
	if err = txn.Set(a.Key("index_cell"), uint64ToBytes(uint64(a.Cell.Index))); err != nil {
		return err
	}
	if err = txn.Set(a.Key("real_cell"), uint64ToBytes(uint64(a.Cell.Real))); err != nil {
		return err
	}
	if err = txn.Set(a.Key("resolution"), uint32ToBytes(a.Cell.Resolution)); err != nil {
		return err
	}
	return txn.Commit()
}

func (db *AlertStore) PutAttr(alert *Alert, attr string, value []byte) error {
	txn := db.NewTransaction(true)
	defer txn.Commit()
	return txn.Set(alert.Key(attr), value)
}

func (db *AlertStore) Delete(a *Alert) error {
	var err error
	txn := db.NewTransaction(true)
	if err = txn.Delete(a.Key("message")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("timestamp")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("read_at")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("delay")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("latitude")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("longitude")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("ephemeral")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("base_cell")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("index_cell")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("real_cell")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("resolution")); err != nil {
		return err
	}
	return txn.Commit()
}

func (db *AlertStore) Size(userID uint32, cellID uint64) uint32 {
	var count uint32
	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := basePrefix(userID, cellID)
	itr.Seek(pre)
	for itr.ValidForPrefix(pre) {
		count++
		itr.Next()
	}
	return count / 7
}

func (db *AlertStore) Exist(a *Alert) bool {
	exist := false
	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := basePrefix(a.UserID, uint64(a.Cell.Index))
	itr.Seek(pre)
	for itr.ValidForPrefix(pre) {
		keySplit := strings.Split(string(itr.Item().Key()), ":")
		if keySplit[len(keySplit)-1] == "message" {
			itr.Item().Value(func(v []byte) error {
				if string(v) == a.Message {
					exist = true
				}
				return nil
			})
		}
		itr.Next()
	}
	return exist
}

func basePrefix(userID uint32, cellID uint64) []byte {
	return []byte(fmt.Sprintf("alert:%d:%d", userID, cellID))
}

// func (db *AlertStore) keyExist(attr string) bool {
// 	txn := db.NewTransaction(false)
// 	itm, err := txn.Get(a.Key(attr))
// 	if err != nil || itm == nil {
// 		return false
// 	}
// 	return true
// }
