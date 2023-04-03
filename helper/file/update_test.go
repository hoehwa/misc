package myfile

import "testing"

func TestUpdateFile(t *testing.T) {
	UpdateFile("example.txt", []byte("Test for UpdateFile method is done successfully!"))
}
