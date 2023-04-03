package myfile

import (
	"os"
)

func UpdateFile(filename string, data []byte) error {
    // Check if the file exists
    _, err := os.Stat(filename)

    // Create the file if it doesn't exist
    if os.IsNotExist(err) {
        _, err := os.Create(filename)
        if err != nil {
					panic(err)
        }
    }

    // Open the file with write-only mode and truncate its contents
    file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
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
