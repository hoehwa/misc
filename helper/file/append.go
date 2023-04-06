package myfile

import (
	"os"
	"sync"
)

// ensure that only one goroutine can access the file at a time
// prevent concurrent writes to the file from causing data corruption.
var mu sync.Mutex

func AppendToFile(path string, filename string, data []byte) error {
	// If file or directory doesn't exist, make one.
	if _, err := os.Stat(path + filename); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		os.Create(path + filename)
	}
	
	mu.Lock()
  defer mu.Unlock()

  file, err := os.OpenFile(path + filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = file.Write(data)
  if err != nil {
    return err
  }

  return nil
}