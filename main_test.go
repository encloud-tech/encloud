package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadContentHandler(t *testing.T) {
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

	var responseObject UploadContentResponse
	json.Unmarshal(data, &responseObject)

	assert.NotNil(t, responseObject.Data)
}

func TestFetchContentHandler(t *testing.T) {
	param := url.Values{}
	param.Set("public_key", "MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA2J/cboAew9qSgtH+i7Sw72hdaOQ0PjRHJGObl0H+xKL+fCYck2ryKwHGkTUCFBuxex9q7xyT4iDsuJke7iVggb/BhueEBT+977BJZYtCfXR2OEPkbx/FnZDh7K6baSIMm1XxGFHlWyG/x5+HpUDUqT7gpjIqoeaoMQ4wkvyKCT6GK6S+6HvTGTmBCkPijMx00+ucQFzozbQWJFXmfwTkz80O89zhA80k0VLeTc2vahMsYpe/JDPa4rWFNU8A+Tcg2aqCK7D8/kIqQw1EQKrl87tph3qBiRuT+U2z5c1VpsAH6eXUJJsYu7/nsSCJ0LeOm070Laf/0H9hwk2HPWIULGq2aCGNhufeHWOh8L7bYrZlJOY3z6BcBng8rS0FzGpyuDkryGB9c5L3JmOIhIyoMETGgn59CRNQgs95pbPx5B+pLCSatC4GJ8nC9seN7RIt3rFRNLDr+AHg42LmiY+yhHIKfpNB8pcTrAoWGvsL8i1WCOvkl5UiQA+ybZDu5ucFAgMBAAE=")

	req := httptest.NewRequest(http.MethodPost, "/fetch-content", strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	FetchContentHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var responseObject FetchContentResponse
	json.Unmarshal(data, &responseObject)

	assert.NotNil(t, responseObject.Data)
}

func TestFetchContentByCIDHandler(t *testing.T) {
	param := url.Values{}
	param.Set("public_key", "MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA2J/cboAew9qSgtH+i7Sw72hdaOQ0PjRHJGObl0H+xKL+fCYck2ryKwHGkTUCFBuxex9q7xyT4iDsuJke7iVggb/BhueEBT+977BJZYtCfXR2OEPkbx/FnZDh7K6baSIMm1XxGFHlWyG/x5+HpUDUqT7gpjIqoeaoMQ4wkvyKCT6GK6S+6HvTGTmBCkPijMx00+ucQFzozbQWJFXmfwTkz80O89zhA80k0VLeTc2vahMsYpe/JDPa4rWFNU8A+Tcg2aqCK7D8/kIqQw1EQKrl87tph3qBiRuT+U2z5c1VpsAH6eXUJJsYu7/nsSCJ0LeOm070Laf/0H9hwk2HPWIULGq2aCGNhufeHWOh8L7bYrZlJOY3z6BcBng8rS0FzGpyuDkryGB9c5L3JmOIhIyoMETGgn59CRNQgs95pbPx5B+pLCSatC4GJ8nC9seN7RIt3rFRNLDr+AHg42LmiY+yhHIKfpNB8pcTrAoWGvsL8i1WCOvkl5UiQA+ybZDu5ucFAgMBAAE=")
	param.Set("private_key", "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")
	param.Set("cid", "bafkreiba4z5c3vluy3blrpxyxcn2r3jjxynlgdfao2hk3thutu52a2l5z4")

	req := httptest.NewRequest(http.MethodPost, "/fetch-content-by-cid", strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	FetchContentByCIDHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	var responseObject FetchByCIDContentResponse
	json.Unmarshal(data, &responseObject)

	assert.NotNil(t, responseObject.Data)
}
