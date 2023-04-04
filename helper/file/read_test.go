package myfile

import (
	"fmt"
	"testing"
)

func TestReadfile(t *testing.T) {
	output := ReadFile("./javascript/","example.txt")
	fmt.Println(output)
}
