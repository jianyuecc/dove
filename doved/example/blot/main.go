package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)


func main() {
	db, err := bolt.Open("doved\\example\\blot\\meta.db", 0600, nil)

	if err != nil {
		log.Fatal("connect meta.db failed. ")
	}

	defer db.Close()


	if err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("kes"));err != nil {
			log.Fatal("create failed ", err)
			return err
		}
		if _, err := tx.CreateBucketIfNotExists([]byte("kec"));err != nil {
			log.Fatal("create failed ", err)
			return err
		}
		return nil
	}); err != nil {
		log.Println("update", err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("kes"))
		b.CreateBucketIfNotExists([]byte("kes_level1"))
		b.CreateBucketIfNotExists([]byte("kes_level2"))
		b.CreateBucketIfNotExists([]byte("kes_level3"))
		b.CreateBucketIfNotExists([]byte("kes_level4"))
		b.CreateBucketIfNotExists([]byte("kes_level5"))
		return nil
	}); err != nil {
		log.Println("update", err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		tx.Bucket([]byte("kes")).Put([]byte("duration"), []byte("30d"))
		b := tx.Bucket([]byte("kes")).Bucket([]byte("kes_level1"))
		b.Put([]byte("region"), []byte("Shanghai"))
		return nil
	}); err != nil {
		log.Println("update", err)
	}


	if err := db.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			b.ForEach(func(k, v []byte) error {
				fmt.Println(string(k), ": " + string(v))

				return nil
			})
			return nil
		})
		return nil
	}); err != nil {
		log.Println("update", err)
	}

	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("kes")).Bucket([]byte("kes_level1"))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), ": "+ string(v))
			return nil
		})
		return nil
	}); err != nil {
		log.Println("update", err)
	}


	// TTLName
	// Duration
	// replication
	// ShardDuration

}

