package pluto

type TransactionalStorage interface {
	Begin(write bool, bucketName []byte) (Bucket, error)
}

type TransactionalBucket interface {
	Bucket
	Commit() error
	Rollback() error
}

type Bucket interface {
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Save(key []byte, value []byte) error
}
