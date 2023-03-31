package http

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	target := "example.com"

	timestamp, respUrl := Request(target)

	fmt.Println(timestamp, respUrl)
}
