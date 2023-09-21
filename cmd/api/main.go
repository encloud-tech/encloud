package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/encloud-tech/encloud/pkg/api"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to encloud API\n")
}

func KeyGenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	cfg, err := api.Fetch()
	if err != nil {
		fmt.Print(err)
	}

	var response types.GenerateKeyPairResponse
	keys, err := api.GenerateKeyPair(cfg.Stat.KekType)
	if err != nil {
		response = types.GenerateKeyPairResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.Keys{},
		}
	} else {
		response = types.GenerateKeyPairResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Keys generated successfully.",
			Data:       keys,
		}
	}

	json.NewEncoder(w).Encode(response)
}

func KeysHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var response types.ListKeysResponse
	keys, err := api.ListKeys()
	if err != nil {
		response = types.ListKeysResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.ListKeys{},
		}
	} else {
		response = types.ListKeysResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Keys fetched successfully.",
			Data:       keys,
		}
	}

	json.NewEncoder(w).Encode(response)
}

func ContentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	r.ParseForm()
	kek := ""
	publicKey := r.FormValue("pubkey")
	readPublicKeyFromPath := r.FormValue("readPubFromPath")

	if readPublicKeyFromPath == "true" {
		kek = thirdparty.ReadKeyFile(publicKey)
	} else {
		kek = publicKey
	}

	var response types.ListContentResponse
	contents := api.List(kek)
	response = types.ListContentResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Content fetched successfully.",
		Data:       contents,
	}

	json.NewEncoder(w).Encode(response)
}

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	kek := ""
	privateKey := ""

	r.ParseForm()
	pubkey := r.FormValue("pubkey")
	privkey := r.FormValue("privkey")
	readPublicKeyFromPath := r.FormValue("readPubFromPath")
	readPrivateKeyFromPath := r.FormValue("readPrivFromPath")
	retrievalFileStoragePath := r.FormValue("storage")
	uuid := r.FormValue("uuid")

	if readPublicKeyFromPath == "true" {
		kek = thirdparty.ReadKeyFile(pubkey)
	} else {
		kek = pubkey
	}

	if readPrivateKeyFromPath == "true" {
		privateKey = thirdparty.ReadKeyFile(privkey)
	} else {
		privateKey = privkey
	}

	var response types.RetrieveByUUIDContentResponse
	fileMetaData, err := api.RetrieveByUUID(uuid, kek, privateKey, retrievalFileStoragePath)
	if err != nil {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileMetadata{},
		}
	} else {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content fetched successfully.",
			Data:       fileMetaData,
		}
	}

	json.NewEncoder(w).Encode(response)
}

func ShareHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	kek := ""
	privateKey := ""

	r.ParseForm()
	pubkey := r.FormValue("pubkey")
	privkey := r.FormValue("privkey")
	readPublicKeyFromPath := r.FormValue("readPubFromPath")
	readPrivateKeyFromPath := r.FormValue("readPrivFromPath")
	email := r.FormValue("email")
	uuid := r.FormValue("uuid")

	if readPublicKeyFromPath == "true" {
		kek = thirdparty.ReadKeyFile(pubkey)
	} else {
		kek = pubkey
	}

	if readPrivateKeyFromPath == "true" {
		privateKey = thirdparty.ReadKeyFile(privkey)
	} else {
		privateKey = privkey
	}

	var response types.RetrieveByUUIDContentResponse
	fileMetaData, err := api.Share(uuid, kek, privateKey, email)
	if err != nil {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileMetadata{},
		}
	} else {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content shared successfully.",
			Data:       fileMetaData,
		}
	}

	json.NewEncoder(w).Encode(response)
}

func SharedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	r.ParseForm()
	decryptedDekPath := r.FormValue("decryptedDekPath")
	dekType := r.FormValue("dekType")
	cid := r.FormValue("cid")
	fileName := r.FormValue("fileName")
	retrievalFileStoragePath := r.FormValue("retrievalFileStoragePath")

	var response types.SharedResponse
	err := api.RetrieveSharedContent(decryptedDekPath, dekType, cid, fileName, retrievalFileStoragePath)
	if err != nil {
		response = types.SharedResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	} else {
		response = types.SharedResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content retrieved successfully.",
		}
	}

	json.NewEncoder(w).Encode(response)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	cfg, err := api.Fetch()
	if err != nil {
		fmt.Print(err.Error())
	}

	kek := ""

	file, handle, err := r.FormFile("file")
	pubkey := r.Form.Get("pubkey")
	readPublicKeyFromPath := r.Form.Get("readPubFromPath")
	dekType := r.Form.Get("type")

	if readPublicKeyFromPath == "true" {
		kek = thirdparty.ReadKeyFile(pubkey)
	} else {
		kek = pubkey
	}

	// Create a temporary file.
	tempFile, err := os.CreateTemp("", "*"+filepath.Ext(handle.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copy the content of the uploaded file to the temporary file.
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response types.UploadContentResponse
	uuid, err := api.Upload(tempFile.Name(), cfg.Stat.KekType, dekType, kek)
	if err != nil {
		response = types.UploadContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.Uuid{},
		}
	} else {
		response = types.UploadContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content uploaded successfully.",
			Data:       types.Uuid{Uuid: uuid},
		}
	}

	json.NewEncoder(w).Encode(response)
}

func RetrieveDecryptedDEKsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	r.ParseForm()
	uuid := r.FormValue("uuid")
	kek := r.FormValue("kek")
	privkey := r.FormValue("privkey")

	uuids := strings.Split(uuid, ",")

	var response types.ListContentResponse
	fileList, err := api.RetrieveDecryptedDEKsWithFileMetadataByUUID(uuids, kek, privkey)
	if err != nil {
		response = types.ListContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileData{},
		}
	} else {
		response = types.ListContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content retrieved successfully.",
			Data:       fileList,
		}
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/keygen", KeyGenHandler).Methods("GET")
	router.HandleFunc("/upload", UploadHandler).Methods("POST")
	router.HandleFunc("/keys", KeysHandler).Methods("GET")
	router.HandleFunc("/contents", ContentsHandler).Methods("POST")
	router.HandleFunc("/retrieve", RetrieveHandler).Methods("POST")
	router.HandleFunc("/share", ShareHandler).Methods("POST")
	router.HandleFunc("/shared", SharedHandler).Methods("POST")
	router.HandleFunc("/decrypted-deks", RetrieveDecryptedDEKsHandler).Methods("POST")

	port := ":9000"
	srv := &http.Server{
		Handler: router,
		Addr:    port,
		// Good practice to enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Printf("Listening on Port: %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
