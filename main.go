package main

import (
	"crypto/rand"
	"encoding/json"
	"filecoin-encrypted-data-storage/cmd"
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"filecoin-encrypted-data-storage/types"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to our API\n")
}

func UploadContentHandler(w http.ResponseWriter, r *http.Request) {
	cfg, _ := config.LoadConf("config.yml")
	estuaryService := service.New(cfg)
	w.Header().Set("Content-Type", "application/json")
	file, handler, err := r.FormFile("data")
	kek := r.Form.Get("public_key")
	timestamp := time.Now().Unix()
	if err != nil {
		response := types.UploadContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.UploadResponse{},
		}
		json.NewEncoder(w).Encode(response)
	}
	defer file.Close()

	//generate a random 32 byte key for AES-256
	dek := make([]byte, 32)
	if _, err := rand.Read(dek); err != nil {
		response := types.UploadContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.UploadResponse{},
		}
		json.NewEncoder(w).Encode(response)
	}

	if _, err := os.Stat("assets"); os.IsNotExist(err) {
		err := os.Mkdir("assets", 0777)
		if err != nil {
			response := types.UploadContentResponse{
				Status:     "fail",
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       types.UploadResponse{},
			}
			json.NewEncoder(w).Encode(response)
		}
	}

	err = thirdparty.EncryptFile(dek, handler.Filename, "assets/encrypted.bin")
	if err != nil {
		response := types.UploadContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.UploadResponse{},
		}
		json.NewEncoder(w).Encode(response)
	}
	content := estuaryService.UploadContent("assets/encrypted.bin")
	if content.CID != "" {
		encryptedDek, err := thirdparty.EncryptWithRSA(dek, thirdparty.GetIdRsaPubFromStr(kek))
		if err != nil {
			response := types.UploadContentResponse{
				Status:     "fail",
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       content,
			}
			json.NewEncoder(w).Encode(response)
		}
		fileData := types.FileMetadata{Timestamp: timestamp, Name: handler.Filename, Size: int(handler.Size), FileType: filepath.Ext(handler.Filename), Dek: encryptedDek, Cid: content.CID}
		service.Store(kek+"-"+content.CID, fileData)
	}

	os.Remove("assets/encrypted.bin")
	w.WriteHeader(http.StatusCreated)
	response := types.UploadContentResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Content uploaded successfully.",
		Data:       content,
	}
	json.NewEncoder(w).Encode(response)
}

func ListContentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	kek := r.FormValue("public_key")
	fileMetaData := service.Fetch(kek)
	w.Header().Set("Content-Type", "application/json")
	os.Remove("assets/downloaded.bin")
	response := types.ListContentResponse{
		Status:     "success",
		StatusCode: http.StatusFound,
		Message:    "Content fetched successfully.",
		Data:       fileMetaData,
	}
	json.NewEncoder(w).Encode(response)
}

func RetrieveContentByCIDHandler(w http.ResponseWriter, r *http.Request) {
	cfg, _ := config.LoadConf("config.yml")
	estuaryService := service.New(cfg)

	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	kek := r.FormValue("public_key")
	privateKey := r.FormValue("private_key")
	cid := r.FormValue("cid")
	fileMetaData := service.FetchByCid(kek + "-" + cid)
	decryptedDek, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
	if err != nil {
		response := types.RetrieveByCIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       fileMetaData,
		}
		json.NewEncoder(w).Encode(response)
	}
	filepath := estuaryService.DownloadContent(fileMetaData.Cid)
	err = thirdparty.DecryptFile(decryptedDek, filepath)
	if err != nil {
		response := types.RetrieveByCIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       fileMetaData,
		}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusCreated)
	os.Remove("assets/downloaded.bin")
	response := types.RetrieveByCIDContentResponse{
		Status:     "success",
		StatusCode: http.StatusFound,
		Message:    "Content fetched successfully.",
		Data:       fileMetaData,
	}
	json.NewEncoder(w).Encode(response)
}

func GenerateKeyPairHandler(w http.ResponseWriter, r *http.Request) {
	thirdparty.InitCrypto()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := types.GenerateKeyPairResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Keys generated successfully.",
		Data:       types.Keys{PublicKey: thirdparty.GetIdRsaPubStr(), PrivateKey: thirdparty.GetIdRsaStr()},
	}
	os.Remove(".keys/.idRsaPub")
	os.Remove(".keys/.idRsa")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/generate-key-pair", GenerateKeyPairHandler).Methods("GET")
	router.HandleFunc("/upload-content", UploadContentHandler).Methods("POST")
	router.HandleFunc("/list-content", ListContentHandler).Methods("POST")
	router.HandleFunc("/fetch-content-by-cid", RetrieveContentByCIDHandler).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice to enforce timeouts for servers you create!
		WriteTimeout: 900 * time.Second,
		ReadTimeout:  900 * time.Second,
	}

	cmd.Execute()
	log.Fatal(srv.ListenAndServe())
}
