package file

import (
	"fmt"
	"os"
)

func Read(filePath string) (fileData []byte, err error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()

	// Read the file
	fileData = make([]byte, fileSize)
	_, err = file.Read(fileData)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return
}

func Create(filePath string, data []byte) error {
	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write the data
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data: %v", err)
	}

	return nil
}

func Delete(filePath string) error {
	return os.Remove(filePath)
}
