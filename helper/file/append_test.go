package myfile

import "testing"

func TestAppendToFile(t *testing.T) {
	AppendToFile("./javascript/", "example.txt", "Append1\n")
	AppendToFile("./javascript/", "example.txt", "Append2\n")
	AppendToFile("./javascript/", "example.txt", "Append3\n")
}