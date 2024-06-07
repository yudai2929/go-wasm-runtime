package main

import (
	"fmt"
	"log"
	"os"
)

const (
	filePath = "add.wasm"
)

func main() {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()

	// Read the file
	fileData := make([]byte, fileSize)
	_, err = file.Read(fileData)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	fmt.Printf("fileData: %v\n", fileData)
}
