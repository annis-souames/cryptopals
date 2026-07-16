package cryptutil

import (
	"encoding/base64"
	"os"
)

// Some file helper methods

func ReadBase64File(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(data))
}
