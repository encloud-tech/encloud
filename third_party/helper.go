package thirdparty

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/mem"
)

func GenerateUuid() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

func GetVirtualMemory() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Total
}

func ReadFile(fileName string) []byte {
	f, _ := os.Open(fileName)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
