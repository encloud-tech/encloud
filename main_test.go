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

func TestUploadRequestHandler(t *testing.T) {
	filePath := "testdata/Provider2InputData.csv"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, _ := os.Open(filePath)
	defer file.Close()
	part, _ := writer.CreateFormFile("data", filePath)
	_, _ = io.Copy(part, file)
	pkw, _ := writer.CreateFormField("public_key")
	pkw.Write([]byte("MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA2J/cboAew9qSgtH+i7Sw72hdaOQ0PjRHJGObl0H+xKL+fCYck2ryKwHGkTUCFBuxex9q7xyT4iDsuJke7iVggb/BhueEBT+977BJZYtCfXR2OEPkbx/FnZDh7K6baSIMm1XxGFHlWyG/x5+HpUDUqT7gpjIqoeaoMQ4wkvyKCT6GK6S+6HvTGTmBCkPijMx00+ucQFzozbQWJFXmfwTkz80O89zhA80k0VLeTc2vahMsYpe/JDPa4rWFNU8A+Tcg2aqCK7D8/kIqQw1EQKrl87tph3qBiRuT+U2z5c1VpsAH6eXUJJsYu7/nsSCJ0LeOm070Laf/0H9hwk2HPWIULGq2aCGNhufeHWOh8L7bYrZlJOY3z6BcBng8rS0FzGpyuDkryGB9c5L3JmOIhIyoMETGgn59CRNQgs95pbPx5B+pLCSatC4GJ8nC9seN7RIt3rFRNLDr+AHg42LmiY+yhHIKfpNB8pcTrAoWGvsL8i1WCOvkl5UiQA+ybZDu5ucFAgMBAAE="))
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
	assert.NotNil(t, string(data))
}
