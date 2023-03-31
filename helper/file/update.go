package file

import (
	"os"
)

func UpdateFile(filename string, content string) {
	// Open the file with read and write permissions
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Seek to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Write the updated content to the file
	_, err = file.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
