package geoalt

import "github.com/dgraph-io/badger"

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
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}
