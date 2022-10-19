package thirdparty

import (
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/mem"
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

func GenerateUuid() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

func GetVirtualMemory() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Total
}
