package read

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/julienbreux/triggy/internal/triggy"
)

// Triggy reads a Triggy file for a given directory
func Triggy(dir string, entrypoint string) (*triggy.Triggy, error) {
	path := filepath.Join(dir, entrypoint)
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf(`triggy: No '%s' file found on '%s'`, entrypoint, dir)
	}
	return readTriggy(path)
}

func readTriggy(file string) (*triggy.Triggy, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var t triggy.Triggy
	return &t, yaml.NewDecoder(f).Decode(&t)
}
