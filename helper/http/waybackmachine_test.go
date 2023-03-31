package http

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	target := "example.com"

	timestamp, respUrl := WayBackRequest(target)

	fmt.Println(timestamp, respUrl)
}
