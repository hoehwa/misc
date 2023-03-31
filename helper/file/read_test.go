package file

import (
	"fmt"
	"testing"
)

func TestReadfile(t *testing.T) {
	output := ReadFile("example.txt")
	fmt.Println(output)
}
