package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateFileName(fileName string) string {
	currentTime := time.Now().Format(time.RFC3339)

	data := fileName + currentTime

	hash := sha256.New()
	hash.Write([]byte(data))

	return hex.EncodeToString(hash.Sum(nil))
}
