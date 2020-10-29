package x

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = randLetters[rand.Intn(len(randLetters))]
	}
	return string(b)
}
