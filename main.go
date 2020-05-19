package main

import (
	"log"
	"fmt"

	"github.com/boltdb/bolt"
	"os"
)

func main() {
	database, err := bolt.Open("reeves.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(database.Path())

	if err := database.Update(func(tx *bolt.Tx) error {
		write, err := tx.CreateBucket([]byte("people"))
		if err != nil {
			return err
		}
		if err := write.Put([]byte("Reeves"), []byte("Akwa")); err != nil {
			return err
		}

		if err := write.Put([]byte("Rufus"), []byte("Emare")); err != nil {
			return err
		}

		if err := write.Put([]byte("Wisdom"), []byte("a software developer")); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Access data from within a read-only transactional block.
	if err := database.Update(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("people")).Get([]byte("Rufus"))
		fmt.Printf("Rufus's last name is %s.\n", v)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	//Delete removes the named bucket
	if err := database.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("people")).Delete([]byte("Rufus"))
	}); err != nil {
		log.Fatalf("error occured due to: %s", err)
	}

	if err := database.Update(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("people")).Get([]byte("Rufus"))
		fmt.Printf("Rufus's last name is %s.\n", v)
		return nil
	}); err != nil {
		log.Fatal(err)
	}


	// Close database to release the file lock.
	if err := database.Close(); err != nil {
		log.Fatal(err)
	}
}