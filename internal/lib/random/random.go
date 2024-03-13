package random

import (
	"math/rand"
	"time"
)

func NewRandomString(i int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	newRune := make([]rune, i)
	for i, _ := range newRune {
		newRune[i] = chars[rnd.Intn(len(chars))]
	}
	return string(newRune)
}
