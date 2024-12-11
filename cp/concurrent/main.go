package main

import (
	"io"
	"os"
	"runtime"
	"sync"
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

	var wg sync.WaitGroup
	sem := make(chan struct{}, runtime.NumCPU())
	startTime := time.Now().UTC()
	wg.Add(len(files))
	for _, file := range files {
		go copyFile(&wg, sem, file.Name())
	}
	wg.Wait()
	took := time.Since(startTime)
	log.Printf("copying files concurrently, %d at a time, took: %v", runtime.NumCPU(), took)
	// clean up
	for _, file := range files {
		os.Remove("files/destination/" + file.Name())
	}
}

func copyFile(wg *sync.WaitGroup, sem chan struct{}, filename string) {
	sem <- struct{}{}
	defer func() {
		<-sem
		wg.Done()
	}()

	srcFile, err := os.Open("files/source/" + filename)
	if err != nil {
		log.Println("os.Open error", err.Error())
		return
	}
	defer srcFile.Close()

	outFile, err := os.Create("files/destination/" + filename)
	if err != nil {
		log.Println("os.Create error", err.Error())
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, srcFile)
	if err != nil {
		log.Println("io.Copy error", err.Error())
		return
	}
}
