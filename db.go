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

//func (f *File) Read(p []byte) (n int, err error) {
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

/*
func (f *File) Read(p []byte) (n int, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte("files")).Get([]byte(f.Md5sum))
		p=make([]byte,len(value))
		copy(p, value)
		return nil
	})
	return len(p), err
}

func (f *File) Write(p []byte) (n int, err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("files")).
			Put([]byte(f.Md5sum), p)
	})
	return len(p), err
}

/*
func (c *Cache) Set(key, value []byte) error {
		err := c.db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(c.requestsBucket)
			if err != nil {
				return err
			}
			err = bucket.Put(key, value)
			if err != nil {
				return err
			}
			return nil
		})

		return err
	}

	func (c *Cache) Get(key []byte) (value []byte, err error) {
	   err = c.db.View(func(tx *bolt.Tx) error {
	      bucket := tx.Bucket(c.requestsBucket)
	      if bucket == nil {
	         return fmt.Errorf("Bucket %q not found!", c.requestsBucket)
	      }

	      var buffer bytes.Buffer
	      buffer.Write(bucket.Get(key))

	      value = buffer.Bytes()
	      return nil
	   })

	   return
	}

*/
