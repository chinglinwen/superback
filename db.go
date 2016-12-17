package main

import (
	"bytes"
	"encoding/gob"
	//"fmt"
	"github.com/boltdb/bolt"
)

type File struct {
	Name     string
	Size     int64
	Time     string
	Path     string
	Location string
	Md5sum   string
	//collectDate time.Time
}

func (f *File) Read() (n int, err error) {
	var value []byte
	err = db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte("files")).Get([]byte(f.Md5sum))
		return nil
	})
	dec := gob.NewDecoder(bytes.NewBuffer(value))
	dec.Decode(&f)
	return len(value), err
}

func (f *File) Write() (n int, err error) {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(f)
	if err != nil {
		return 0, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("files")).
			Put([]byte(f.Md5sum), buf.Bytes())
	})
	return buf.Len(), err
}

