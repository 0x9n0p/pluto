package storage

import (
	"fmt"
	"path/filepath"
	"pluto"

	"github.com/tidwall/buntdb"
)

type BuntDB struct {
	Path string
	*buntdb.DB
}

func NewBuntDB(path string) (*BuntDB, error) {
	db, err := buntdb.Open(filepath.Join(pluto.Env.RootStoragePath, path))
	if err != nil {
		return nil, fmt.Errorf("open buntdb: %v", err)
	}

	return &BuntDB{
		Path: path,
		DB:   db,
	}, nil
}

func (b *BuntDB) Begin(write bool, bucketName []byte) (pluto.Bucket, error) {
	return nil, nil
}
