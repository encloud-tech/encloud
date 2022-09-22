package main

import (
	"bytes"
	"encoding/json"
	"filecoin-encrypted-data-storage/types"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PublicKey string
var PrivateKey string
var Cid string

func TestGenerateKeyPairHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/generate-key-pair", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	GenerateKeyPairHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var responseObject types.GenerateKeyPairResponse
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)
	PublicKey = responseObject.Data.PublicKey
	PrivateKey = responseObject.Data.PrivateKey
	assert.NotNil(t, responseObject.Data)
}

func TestUploadContentHandler(t *testing.T) {
	log.Println(PublicKey)
	filePath := "testdata/Provider2InputData.csv"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, _ := os.Open(filePath)
	defer file.Close()
	part, _ := writer.CreateFormFile("data", filePath)
	_, _ = io.Copy(part, file)
	pkw, _ := writer.CreateFormField("public_key")
	pkw.Write([]byte(PublicKey))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload-content", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	UploadContentHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var responseObject types.UploadContentResponse
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)
	Cid = responseObject.Data.CID
	assert.NotNil(t, responseObject.Data)
}

func TestListContentHandler(t *testing.T) {
	log.Println(PublicKey)
	param := url.Values{}
	param.Set("public_key", PublicKey)

	req := httptest.NewRequest(http.MethodPost, "/fetch-content", strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ListContentHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var responseObject types.ListContentResponse
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)
	assert.NotNil(t, responseObject.Data)
}

func TestRetrieveContentByCIDHandler(t *testing.T) {
	log.Println(PublicKey)
	log.Println(PrivateKey)
	log.Println(Cid)
	param := url.Values{}
	param.Set("public_key", PublicKey)
	param.Set("private_key", PrivateKey)
	param.Set("cid", Cid)

	req := httptest.NewRequest(http.MethodPost, "/fetch-content-by-cid", strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	RetrieveContentByCIDHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var responseObject types.RetrieveByCIDContentResponse
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)
	assert.NotNil(t, responseObject.Data)
}
