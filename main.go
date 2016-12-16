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

	//for {
	//}
}

/*

get all the file info

store local
have a ui for it,
  fuzzy search?

compare one by one, if no backup, backup then(upload to somewhere)
can specify multiple path (using parameter)


Remove the duplicate backup of all my files ( store only one in the cloud )
To store the meta info in the cloud as a service

Try run all things to collect that info

Backup only one of them ( program do the backup )

Focus on pdf first? (and other small file)


Super backup
Book index (multi dimension search, of just fuzzy search )
Also provide a exact search

Upload to one location may lost the organized info(that I spend long time on it)

How to make it easier accessable
To modify this meta info
To easier download

It may need a ui to do this

Non pdf file out into other directory
Will the file number exceed to max number in a directory?

Backup to cloud and removable disk, tar from cloud, sync to local, how to do incremental sync

There may have lots operation to list them all
To sync the meta way, light way and heavy way (download to local first, upload the modified one?) So it involves meta update

*/
