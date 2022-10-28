package thirdparty

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"encloud/config"
	"io/ioutil"
	"log"
	"os"
)

func InitCrypto() {
	if _, err := os.Stat(config.DotKeys); err != nil {
		// create it
		err = os.Mkdir(config.DotKeys, 0700)
		if err != nil {
			log.Fatalf("ERROR: fail to create keys dir, %s", err.Error())
			os.Exit(1)
		}
	}

	if _, err := os.Stat(config.IdRsa); err != nil {
		// generate key pair
		// save private key
		// save public key
		keyPair := generateKeyPair()
		saveIdRsa(config.IdRsa, keyPair)
		saveIdRsaPub(config.IdRsaPub, keyPair)
	}
}

func generateKeyPair() *rsa.PrivateKey {
	// generate key pair
	keyPair, err := rsa.GenerateKey(rand.Reader, config.KeySize)
	if err != nil {
		log.Printf("ERROR: fail to create key pair, %s", err.Error())
		os.Exit(1)
	}

	// validate key
	err = keyPair.Validate()
	if err != nil {
		log.Printf("ERROR: fail to validate key pair, %s", err.Error())
		os.Exit(1)
	}

	return keyPair
}

func saveIdRsa(fileName string, keyPair *rsa.PrivateKey) {
	// private key stream
	privateKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(keyPair),
	}

	log.Print(privateKeyBlock)

	// create file
	f, err := os.Create(fileName)
	if err != nil {
		log.Printf("ERROR: fail to save idrsa, %s", err.Error())
		os.Exit(1)
	}

	err = pem.Encode(f, privateKeyBlock)
	if err != nil {
		log.Printf("ERROR: fail to save idrsa, %s", err.Error())
		os.Exit(1)
	}
}

func saveIdRsaPub(fileName string, keyPair *rsa.PrivateKey) {
	// public key stream
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&keyPair.PublicKey)
	if err != nil {
		log.Printf("ERROR: fail to save idrsapub, %s", err.Error())
		os.Exit(1)
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}

	// create file
	f, err := os.Create(fileName)
	if err != nil {
		log.Printf("ERROR: fail to save idrsapub, %s", err.Error())
		os.Exit(1)
	}

	err = pem.Encode(f, publicKeyBlock)
	if err != nil {
		log.Printf("ERROR: fail to save idrsapub, %s", err.Error())
		os.Exit(1)
	}
}

func GetIdRsa() *rsa.PrivateKey {
	keyData, err := ioutil.ReadFile(config.IdRsa)
	if err != nil {
		log.Printf("ERROR: fail get idrsa, %s", err.Error())
		os.Exit(1)
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsa, invalid key")
		os.Exit(1)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		log.Printf("ERROR: fail get idrsa, %s", err.Error())
		os.Exit(1)
	}

	return privateKey
}

func GetIdRsaPub() *rsa.PublicKey {
	keyData, err := ioutil.ReadFile(config.IdRsaPub)
	if err != nil {
		log.Printf("ERROR: fail get idrsapub, %s", err.Error())
		return nil
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsapub, invalid key")
		return nil
	}

	publicKey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
	if err != nil {
		log.Printf("ERROR: fail get idrsapub, %s", err.Error())
		return nil
	}
	switch publicKey := publicKey.(type) {
	case *rsa.PublicKey:
		return publicKey
	default:
		return nil
	}
}

func GetIdRsaStr() string {
	keyData, err := ioutil.ReadFile(config.IdRsa)
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

func GetIdRsaPubStr() string {
	keyData, err := ioutil.ReadFile(config.IdRsaPub)
	if err != nil {
		log.Printf("ERROR: fail get idrsapubstr, %s", err.Error())
		return ""
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsapubstr, invalid key")
		return ""
	}

	// encode base64 key data
	return base64.StdEncoding.EncodeToString(keyBlock.Bytes)
}

func GetIdRsaFromStr(keyStr string) *rsa.PrivateKey {
	// key is base64 encoded
	data, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		log.Printf("ERROR: fail get rsa, %s", err.Error())
		return nil
	}

	// get rsa private key
	key, err := x509.ParsePKCS1PrivateKey(data)
	if err != nil {
		log.Printf("ERROR: fail get rsa, %s", err.Error())
		return nil
	}
	// switch key := key.(type) {
	// case *rsa.PrivateKey:
	return key
	// default:
	// return nil
	// }
}

func GetIdRsaPubFromStr(keyStr string) *rsa.PublicKey {
	// key is base64 encoded
	data, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		log.Printf("ERROR: fail get rsapub, %s", err.Error())
		return nil
	}

	// this for ios key
	var pubKey rsa.PublicKey
	if rest, err := asn1.Unmarshal(data, &pubKey); err != nil {
		log.Println("INFO: not ios key", keyStr)
	} else if len(rest) != 0 {
		log.Printf("INFO: not ios key, invalid lenght, %s", keyStr)
	} else {
		return &pubKey
	}

	// this is for android
	// get rsa public key
	pub, err := x509.ParsePKIXPublicKey(data)
	if err != nil {
		log.Printf("INFO: not android key, %s", keyStr)
		return nil
	}
	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub
	default:
		return nil
	}
}

func EncryptWithRSA(payload []byte, key *rsa.PublicKey) ([]byte, error) {
	rnd := rand.Reader
	hash := sha256.New()

	// encrypt with OAEP
	ciperText, err := rsa.EncryptOAEP(hash, rnd, key, payload, nil)
	if err != nil {
		log.Printf("ERROR: fail to encrypt, %s", err.Error())
		return nil, err
	}

	return ciperText, nil
}

func DecryptWithRSA(payload []byte, key *rsa.PrivateKey) ([]byte, error) {
	rnd := rand.Reader
	hash := sha256.New()

	// decrypt with OAEP
	plainText, err := rsa.DecryptOAEP(hash, rnd, key, payload, nil)
	if err != nil {
		log.Printf("ERROR: fail to decrypt, %s", err.Error())
		return nil, err
	}

	return plainText, nil
}
