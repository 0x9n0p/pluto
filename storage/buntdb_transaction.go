package storage

import "github.com/tidwall/buntdb"

type BuntDBTransaction struct {
	*buntdb.Tx
}

func (b *BuntDBTransaction) Get(key []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BuntDBTransaction) Delete(key []byte) error {
	//TODO implement me
	panic("implement me")
}

func (b *BuntDBTransaction) Save(key []byte, value []byte) error {
	//TODO implement me
	panic("implement me")
}
