package object_name

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	length  = 12
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var seedRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}

	return string(b)
}

func GenerateUniqueObjectName(filename string) string {
	timestamp := time.Now().UnixNano()
	randomStr := generateRandomString(length)
	objName := fmt.Sprintf("%d-%s-%s", timestamp, randomStr, filename)

	return objName
}
