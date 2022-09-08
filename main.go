package main

import (
	"crypto/rand"
	"encoding/json"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type Keys struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to our API\n")
}

func UploadContentHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("data")
	kek := r.Form.Get("public_key")
	timestamp := time.Now().Unix()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//generate a random 32 byte key for AES-256
	dek := make([]byte, 32)
	if _, err := rand.Read(dek); err != nil {
		panic(err.Error())
	}

	if _, err := os.Stat("assets"); os.IsNotExist(err) {
		err := os.Mkdir("assets", 0777)
		if err != nil {
			panic(err)
		}
	}

	thirdparty.EncryptFile(dek, file)
	content := service.UploadContent("assets/encrypted.bin")
	if content.CID != "" {
		encryptedDek, err := thirdparty.EncryptWithRSA(dek, thirdparty.GetIdRsaPubFromStr(kek))
		if err != nil {
			panic(err.Error())
		}
		fileData := service.FileMetadata{Timestamp: timestamp, Name: handler.Filename, Size: int(handler.Size), FileType: filepath.Ext(handler.Filename), Dek: encryptedDek, Cid: content.CID}
		service.Store(kek, fileData)
	}

	os.Remove("assets/encrypted.bin")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(content)
}

func FetchContentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	kek := r.FormValue("public_key")
	privateKey := r.FormValue("private_key")
	fileMetaData := service.Fetch(kek)
	decryptedDek, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
	if err != nil {
		panic(err.Error())
	}
	filepath := service.DownloadContent(fileMetaData.Cid)
	thirdparty.DecryptFile(decryptedDek, filepath)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	os.Remove("assets/downloaded.bin")
	json.NewEncoder(w).Encode(fileMetaData)
}

func GenerateKeyPairHandler(w http.ResponseWriter, r *http.Request) {
	thirdparty.InitCrypto()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := Keys{PublicKey: thirdparty.GetIdRsaPubStr(), PrivateKey: thirdparty.GetIdRsaStr()}
	os.Remove(".keys/.idRsaPub")
	os.Remove(".keys/.idRsa")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/generate-key-pair", GenerateKeyPairHandler).Methods("GET")
	router.HandleFunc("/upload-content", UploadContentHandler).Methods("POST")
	router.HandleFunc("/fetch-content", FetchContentHandler).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9000",
		// Good practice to enforce timeouts for servers you create!
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
