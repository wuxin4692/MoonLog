package main

import (
	"os"
	"fmt"
	"time"
)

var StartSize int64


func GetFileSize(filename string) int64 {
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	size,err := file.Seek(0, os.SEEK_END)
	if err != nil {
		fmt.Println(err)
	}
	return size
}



func main() {

	FileName := ReadConf("moonlog.yaml").LogPath
	StartSize = GetFileSize(FileName)
	for {
		time.Sleep(1 * time.Second)
		if StartSize == GetFileSize(FileName) {
			continue
		}
		file,err := os.Open(FileName)
		defer file.Close()
		if err != nil {
			return
		}
		Data := make([]byte, StartSize)
		file.ReadAt(Data,StartSize)
		fmt.Println(string(Data[:]))
		StartSize = GetFileSize(FileName)
	}
}