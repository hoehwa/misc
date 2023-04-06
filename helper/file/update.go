package myfile

import (
	"os"
)

func UpdateFile(path string,filename string, data []byte) error {
    // If file or directory doesn't exist, make one.
    if _, err := os.Stat(path + filename); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		os.Create(path + filename)
	}
    
    // Open the file with write-only mode and truncate its contents
    file, err := os.OpenFile(path + filename, os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
		panic(err)
    }
    defer file.Close()

    // Write the data to the file
    _, err = file.Write(data)
    if err != nil {
		panic(err)
    }

    return nil
}
