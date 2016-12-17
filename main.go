package main

import (
	//"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

type Paths struct {
	Path []string
}

type Files struct {
	Path  []string
	Files []File
}

/*

//file or dir
func process (ch chan string) {
	for {
		select {
			case f:= <-ch :

		}

	}
}
*/

var db *bolt.DB

const timelayout = "20060102_150405"

func main() {
	var err error
	db, err = bolt.Open("bolt.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("files"))
		return err
	})

	pwd, _ := os.Getwd()
	err = scanDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

}

