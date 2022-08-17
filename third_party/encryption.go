package thirdparty

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
)

func EncryptFile(key []byte, file multipart.File) {
	// Reading file
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, buf.Bytes(), nil)

	// Writing ciphertext file
	err = ioutil.WriteFile("assets/encrypted.bin", cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

}

func DecryptFile(key []byte) {
	// Reading encrypted file
	cipherText, err := ioutil.ReadFile("assets/encrypted.bin")
	if err != nil {
		log.Fatal(err)
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	// Writing decryption content
	err = ioutil.WriteFile("assets/plaintext.txt", plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
