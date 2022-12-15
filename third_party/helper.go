package thirdparty

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/pem"
	"fmt"
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

func ReadKeyFile(filePath string) string {
	f, _ := os.Open(filePath)
	keyData, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("ERROR: fail get idrsa, %s", err.Error())
		os.Exit(1)
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsa, invalid key")
		os.Exit(1)
	}

	// encode base64 key data
	return base64.StdEncoding.EncodeToString(keyBlock.Bytes)
}

func DigestString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
