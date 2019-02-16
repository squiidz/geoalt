package geoalt

import (
	"runtime"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/badger/options"
)

type DB struct {
	UserStore  *UserStore
	AlertStore *AlertStore
}

func NewDB(userPath, alertPath string) (*DB, error) {
	userdb, err := openBadger(userPath)
	if err != nil {
		return nil, err
	}
	alertdb, err := openBadger(alertPath)
	if err != nil {
		return nil, err
	}
	return &DB{
		UserStore:  &UserStore{userdb},
		AlertStore: &AlertStore{alertdb},
	}, nil
}

func openBadger(path string) (*badger.DB, error) {
	opts := badger.DefaultOptions
	if runtime.GOARCH == "arm" {
		opts.ValueLogLoadingMode = options.FileIO
	}
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}
