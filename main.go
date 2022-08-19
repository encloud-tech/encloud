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
	"time"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to our API\n")
}

func uploadContent(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	thirdparty.EncryptFile(bytes, file)
	content := service.UploadContent("assets/encrypted.bin")
	os.Remove("assets/encrypted.bin")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(content)
}

func fetchContent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	content := service.FetchContentByCid(id)
	json.NewEncoder(w).Encode(content)
}

func fetchAllContents(w http.ResponseWriter, r *http.Request) {
	contents := service.FetchAllContents()
	json.NewEncoder(w).Encode(contents)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contents", fetchAllContents).Methods("GET")
	router.HandleFunc("/contents/{id}", fetchContent).Methods("GET")
	router.HandleFunc("/content", uploadContent).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9000",
		// Good practice to enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
