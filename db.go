package geoalt

import "github.com/dgraph-io/badger"

type DB struct {
	*badger.DB
}

func NewDB(path string) (*DB, error) {
	return openBadger(path)
}

func openBadger(path string) (*DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
