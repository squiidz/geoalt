package geoalt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dgraph-io/badger"
	h3 "github.com/uber/h3-go"
)

type AlertStore struct {
	*badger.DB
}

type Alert struct {
	ID        uint32
	CellID    h3.H3Index // Resolution of indexing
	MinCell   h3.H3Index // Smallest resultion
	CellRes   uint32     // Resolution for geo fence
	Coord     Coord
	UserID    uint32
	Message   string
	Timestamp string
	Ephemeral bool
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
	return []byte(fmt.Sprintf("alert:%d:%d:%d:%s", a.UserID, a.CellID, a.ID, attr))
}

func (a *Alert) SetAttr(attr string, value []byte) {
	switch attr {
	case "message":
		a.Message = string(value)
	case "timestamp":
		a.Timestamp = string(value)
	case "latitude":
		a.Coord.Lat = float64fromBytes(value)
	case "longitude":
		a.Coord.Lng = float64fromBytes(value)
	case "ephemeral":
		a.Ephemeral = boolFromBytes(value)
	case "min_cell":
		a.MinCell = h3.H3Index(uint64FromBytes(value))
	case "cell_res":
		a.CellRes = uint32FromBytes(value)
	}
}

func (a *Alert) Borders() []*Coord {
	var coords []*Coord
	rcell := h3.ToParent(a.MinCell, int(a.CellRes))
	boundaries := h3.ToGeoBoundary(rcell)
	for _, b := range boundaries {
		coords = append(coords, &Coord{
			Lat: b.Latitude,
			Lng: b.Longitude,
		})
	}
	return coords
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
		alert.UserID = userID
		itr.Next()
	}
	alert.ID = id
	alert.CellID = h3.H3Index(cellID)
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
	a.ID = db.Size(a.UserID, uint64(a.CellID)) + 1
	a.MinCell = h3.FromGeo(h3.GeoCoord{Latitude: a.Coord.Lat, Longitude: a.Coord.Lng}, 15)
	if err = txn.Set(a.Key("message"), []byte(a.Message)); err != nil {
		return err
	}
	if err = txn.Set(a.Key("timestamp"), []byte(a.Timestamp)); err != nil {
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
	if err = txn.Set(a.Key("min_cell"), uint64ToBytes(uint64(a.MinCell))); err != nil {
		return err
	}
	if err = txn.Set(a.Key("cell_res"), uint32ToBytes(a.CellRes)); err != nil {
		return err
	}
	return txn.Commit()
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
	if err = txn.Delete(a.Key("latitude")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("longitude")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("ephemeral")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("min_cell")); err != nil {
		return err
	}
	if err = txn.Delete(a.Key("cell_res")); err != nil {
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
	return count
}

func (db *AlertStore) Exist(a *Alert) bool {
	exist := false
	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	pre := basePrefix(a.UserID, uint64(a.CellID))
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
