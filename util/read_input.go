package util

import (
	"os"
)

func Read(day, fileName string) ([]byte, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(path + "/" + day + "/" + fileName)
	if err != nil {
		return nil, err
	}

	return b, nil
}
