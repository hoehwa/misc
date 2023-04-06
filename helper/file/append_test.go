package myfile

import "testing"

func TestAppendToFile(t *testing.T) {
	AppendToFile("./javascript/", "example.txt", []byte("Append1\n"))
	AppendToFile("./javascript/", "example.txt", []byte("Append2\n"))
	AppendToFile("./javascript/", "example.txt", []byte("Append3\n"))
}