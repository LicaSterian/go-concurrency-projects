package main

import (
	"io"
	"os"
	"time"

	"log"
)

func main() {
	// list files of current directory
	files, err := os.ReadDir("files/source/")
	if err != nil {
		log.Println("os.ReadDir error", err.Error())
		return
	}

	startTime := time.Now().UTC()

	for _, file := range files {
		srcFile, err := os.Open("files/source/" + file.Name())
		if err != nil {
			log.Println("os.Open error", err.Error())
			return
		}

		outFile, err := os.Create("files/destination/" + file.Name())
		if err != nil {
			log.Println("os.Create error", err.Error())
			return
		}

		_, err = io.Copy(outFile, srcFile)
		if err != nil {
			log.Println("io.Copy error", err.Error())
			return
		}
		srcFile.Close()
		outFile.Close()
	}

	took := time.Since(startTime)
	log.Printf("copying files one-by-one took: %v", took)
	// clean up
	for _, file := range files {
		os.Remove("files/destination/" + file.Name())
	}
}
