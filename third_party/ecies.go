package thirdparty

import (
	"log"

	ecies "github.com/ecies/go/v2"
)

func GenerateEciesKeyPair() {
	keyPair, err := ecies.GenerateKey()
	if err != nil {
		panic(err)
	}
	log.Println("key pair has been generated")
	log.Println(keyPair)
}
