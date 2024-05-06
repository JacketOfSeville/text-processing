package utils

import (
	"errors"
	"os"
)

func ReadFileFromPath(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return nil, err
	}

	// Kind of an early optimization, should not be faster for small files
	// but should be a lot faster for large files
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		return nil, err
	}

	return bs, nil
}

func ExtractExtension(filename string) (string, error) {
	// find the last dot
	lastDotIndex := -1
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			lastDotIndex = i
			break
		}
	}

	if lastDotIndex == -1 {
		return "", errors.New("No extension found")
	}

	return filename[lastDotIndex:], nil
}
