package mitre

import (
	"hash"
	"io"
	"os"
)

func hashSum(filePath string, hashFunction func() hash.Hash) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte{0}, err
	}
	defer file.Close()

	hash := hashFunction()
	if _, err := io.Copy(hash, file); err != nil {
		return []byte{0}, err
	}
	return hash.Sum(nil), nil
}
