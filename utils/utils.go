package utils

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GetRandomStr(secure bool, length int, charsets ...[]rune) string {
	var letters []rune
	if len(charsets) == 0 {
		letters = defaultLetters
	} else {
		letters = charsets[0]
	}

	b := make([]rune, length)

	for i := range b {
		if secure {
			n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			b[i] = letters[n.Int64()]
		} else {
			mrand.Seed(time.Now().UnixNano())
			b[i] = letters[mrand.Intn(len(letters))]
		}

	}
	return string(b)
}
