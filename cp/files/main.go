package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// generate 1000 files and write only one byte to each of them
	for i := 0; i < 1000; i++ {
		// create file
		filename := fmt.Sprintf("source/file_%d", i)
		file, err := os.Create(filename)
		if err != nil {
			log.Println("os.Create error", err.Error())
			return
		}

		// write zero byte
		_, err = file.Write([]byte{1})
		if err != nil {
			log.Println("file.Write error", err.Error())
			return
		}

		// close file
		file.Close()
	}
}
