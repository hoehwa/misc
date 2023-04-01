package myos

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindRootDir() (string, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return "", err
    }

    // Traverse up the directory hierarchy until we find a directory that
    // contains a "go.mod" file, indicating that we've found the root directory.
    for {
        if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
            return cwd, nil
        }

        parent := filepath.Dir(cwd)
        if parent == cwd {
            return "", fmt.Errorf("could not find root directory")
        }

        cwd = parent
    }
}