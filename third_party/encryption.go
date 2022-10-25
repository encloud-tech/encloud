package thirdparty

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func EncryptFile(key []byte, filePath string, uploadPath string) error {
	const BufferSize = 200 * 1000000 // 200 MB

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		// Creating block of algorithm
		block, err := aes.NewCipher(key)
		if err != nil {
			log.Fatalf("cipher err: %v", err.Error())
			return err
		}

		// Creating GCM mode
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			log.Fatalf("cipher GCM err: %v", err.Error())
			return err
		}

		// Generating random nonce
		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			log.Fatalf("nonce  err: %v", err.Error())
			return err
		}

		// Decrypt file
		cipherText := gcm.Seal(nonce, nonce, buffer[:bytesread], nil)

		f, err := os.OpenFile(uploadPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.Write(cipherText); err != nil {
			panic(err)
		}

		fmt.Println("bytes read: ", bytesread)
	}

	return nil
}

func DecryptFile(key []byte, filepath string) error {
	// Reading encrypted file
	cipherText, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
		return err
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
		return err
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
		return err
	}

	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
		return err
	}

	// Writing decryption content
	err = ioutil.WriteFile("assets/decrypted.txt", plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
		return err
	}

	return nil
}
