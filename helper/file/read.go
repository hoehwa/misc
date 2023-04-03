package myfile

import (
	"os"
)

func ReadFile(filename string) string {
	// If file does not exist, make a file.
	file, err := os.Open(filename)
	if err != nil {
		os.Create(filename)
	}
	defer file.Close()

	// Read the file contents into a byte slice
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}

	// Convert the byte slice to a string and print it
	fileContent := string(buffer)
	return fileContent
}
