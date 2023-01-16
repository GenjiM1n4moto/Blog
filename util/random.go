package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(9))
}

func RandomPassword() string {
	return RandomString(6)
}
func RandomNickname() string {
	return RandomString(6)
}
