package myfile

import (
	"os"
)

func ReadFile(path string, filename string) string {
	// If file or directory doesn't exist, make one.
	if _, err := os.Stat(path + filename); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		os.Create(path + filename)
	}

	file, err := os.Open(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file contents into a byte slice
	fileInfo, _ := file.Stat()
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
