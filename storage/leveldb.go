package storage

import "github.com/syndtr/goleveldb/leveldb"

type LevelDBAdapter struct {
	backend *leveldb.DB
}

func NewLevelDBAdapter(path string) (*LevelDBAdapter, error) {
	backend, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBAdapter{backend}, nil
}

func (db *LevelDBAdapter) Get(key []byte) ([]byte, error) {
	return db.backend.Get(key, nil)
}

func (db *LevelDBAdapter) Put(key, value []byte) error {
	return db.backend.Put(key, value, nil)
}

func (db *LevelDBAdapter) Has(key []byte) bool {
	has, _ := db.backend.Has(key, nil)
	return has
}

func (db *LevelDBAdapter) Delete(key []byte) error {
	return db.backend.Delete(key, nil)
}

func (db *LevelDBAdapter) BatchPut(kvs [][2][]byte) error {
	batch := new(leveldb.Batch)
	for i := range kvs {
		batch.Put(kvs[i][0], kvs[i][1])
	}
	return db.backend.Write(batch, nil)
}

func (db *LevelDBAdapter) Close() {
	db.backend.Close()
}
