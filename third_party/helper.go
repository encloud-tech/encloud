package thirdparty

import (
	"math/rand"
	"strings"
)

func RandomString(n int, alphabet []rune) string {
	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
