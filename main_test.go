package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestHandler(t *testing.T) {
	filePath := "testdata/Provider2InputData.csv"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, _ := os.Open(filePath)
	defer file.Close()
	part, _ := writer.CreateFormFile("data", filePath)
	_, _ = io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/content", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	UploadContentHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, string(data))
}
