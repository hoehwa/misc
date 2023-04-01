package os

import (
	"fmt"
	"testing"
)

func TestFindRootDir(t *testing.T) {
	fmt.Println(FindRootDir())
}