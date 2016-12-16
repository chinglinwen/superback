package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func getmd5(filename string, size int64) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		//return "", fmt.Errorf("%v,%#v",filename,err)
		return "", err
	}

	var length int64
	if size > 100 {
		length = 100
	} else {
		length = size
	}

	buf := make([]byte, length)
	_, err = f.Read(buf)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5.Sum(buf)), nil
}
