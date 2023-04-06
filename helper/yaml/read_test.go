package yaml

import (
	"fmt"
	"testing"
)

func TestReadOneDemensionalYaml(t *testing.T) {
	fmt.Println(ReadOneDemensionalYaml("./simple/", "one-demension.yaml"))
}