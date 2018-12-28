package gorand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandStrWithCharset returns a random string of length as mentioned in the argument, created
// out of characters in the charset string
func RandStrWithCharset(length int, charset string) string {
	rand := make([]byte, length)
	for i := range rand {
		rand[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(rand)
}

func RandStr(length int) string {
	return RandStrWithCharset(length, charset)
}
