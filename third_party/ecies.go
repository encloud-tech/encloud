package thirdparty

import (
	"log"

	ecies "github.com/ecies/go/v2"
)

func EciesGenerateKeyPair() *ecies.PrivateKey {
	keyPair, err := ecies.GenerateKey()
	if err != nil {
		panic(err)
	}

	return keyPair
}

func EncryptWithEcies(pubkey *ecies.PublicKey, msg []byte) ([]byte, error) {
	ciphertext, err := ecies.Encrypt(pubkey, msg)
	if err != nil {
		log.Printf("ERROR: fail to encrypt, %s", err.Error())
		return nil, err
	}

	return ciphertext, nil
}

func DecryptWithEcies(privkey *ecies.PrivateKey, msg []byte) ([]byte, error) {
	plaintext, err := ecies.Decrypt(privkey, msg)
	if err != nil {
		log.Printf("ERROR: fail to decrypt, %s", err.Error())
		return nil, err
	}

	return plaintext, nil
}

func NewPrivateKeyFromHex(s string) *ecies.PrivateKey {
	pub, err := ecies.NewPrivateKeyFromHex(s)
	if err != nil {
		log.Printf("ERROR: fail to convert, %s", err.Error())
	}

	return pub
}

func NewPublicKeyFromHex(s string) *ecies.PublicKey {
	pub, err := ecies.NewPublicKeyFromHex(s)
	if err != nil {
		log.Printf("ERROR: fail to convert, %s", err.Error())
	}

	return pub
}
