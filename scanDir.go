package main

import (
	"fmt"
	"path"
	"strings"
	//"time"
	"io/ioutil"
)

func scanDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		filename := f.Name()
		size := f.Size()
		time := f.ModTime().Format(timelayout)

		if f.IsDir() {
			scanDir(path.Join(dir, filename))
			continue
		}

		if !strings.HasSuffix(strings.ToLower(filename), "pdf") {
			continue
		}

		md5, err := getmd5(path.Join(dir, filename), size)
		if err != nil {
			fmt.Printf("getmd5 %v ,%v, %v\n", dir, filename, err)
			continue
		}

		data := File{
			Name:     filename,
			Size:     size,
			Time:     time,
			Path:     dir,
			Location: "home",
			Md5sum:   md5,
		}

		_, err = data.Write()
		if err != nil {
			fmt.Printf("write %v ,%v, %v\n", dir, filename, err)
		}

		data2 := File{Md5sum: md5}
		data2.Read()
		fmt.Printf("data2 %#v\n", data2)
	}
	return nil
}
