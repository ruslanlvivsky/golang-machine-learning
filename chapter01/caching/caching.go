package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func main() {
	// 메모리 캐싱
	c := cache.New(5*time.Minute, 30*time.Second)
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}

	// 디스크 로컬 캐싱
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err = b.Put([]byte("mykey"), []byte("myvalue"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		cursor := b.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			fmt.Printf("key: %s, value: %s\n", key, value)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
