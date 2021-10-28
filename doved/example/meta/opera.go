package meta

import (
	"github.com/boltdb/bolt"
	"log"
	"sync"
)

type Meta struct {
	DB *bolt.DB
	Mutex *sync.Mutex
	FileName string

}



func (m *Meta) OpenDB()  {
	var err error
	m.DB, err = bolt.Open(m.FileName, 0600, nil)

	if err != nil {
		log.Fatalln("open the boltdb failed: ",err)
	}
}

func (m Meta) Put(key []byte, value []byte)  {
	m.OpenDB()
	defer m.DB.Close()
	m.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("system"))
		bucket.Put(key, value)
		return nil
	})
}

func (m Meta) PutCPU(key []byte, value []byte)  {
	m.OpenDB()
	defer m.DB.Close()
	m.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("system")).Bucket([]byte("cpu"))
		bucket.Put(key, value)
		return nil
	})
}
func (m *Meta) ShowDatabases() (error,[]string) {
	m.OpenDB()
	defer m.DB.Close()
	var databases []string
	return m.DB.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			databases = append(databases, string(name))
			return nil
		})
	}), databases

}

func (m Meta) ShowTables(database string) (error,[]string) {
	m.OpenDB()
	defer m.DB.Close()
	var tables []string

	return m.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database))
		b.ForEach(func(k, v []byte) error {
			tables = append(tables, string(k))
			return nil
		})
		return nil

	}), tables
}

func (m *Meta) CreateDatabase(database string) error {
	m.OpenDB()
	defer m.DB.Close()
	return m.DB.Update(func(tx *bolt.Tx) error {
		_,err :=  tx.CreateBucketIfNotExists([]byte(database))
		return  err
	})
}

func (m *Meta) CreateTable(database string, table string) error {
	m.OpenDB()
	defer m.DB.Close()
	return m.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.Bucket([]byte(database)).CreateBucketIfNotExists([]byte(table))
		return err
	})
}


