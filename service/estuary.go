package service

import (
	"bytes"
	"encoding/json"
	"filecoin-encrypted-data-storage/constants"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Contents []Content
type ByCID []struct {
	Content ByCidResponse
}

type Content struct {
	Name string `json:"name"`
	CID  string `json:"cid"`
}

type UploadResponse struct {
	CID       string
	EstuaryId int
}

type ByCidResponse struct {
	Name string
	CID  string
}

func FetchAllContents() Contents {
	log.Print("Start fetching data from service")
	response := doApiRequest(
		"GET",
		constants.EstuaryApiBaseUrl+"/content/list",
		nil,
	)
	var responseObject Contents
	json.Unmarshal(response, &responseObject)
	log.Print("Data fetched from service: ", responseObject)
	return responseObject
}

func FetchContentByCid(cid string) ByCID {
	log.Print("Start fetching data from service")
	response := doApiRequest(
		"GET",
		constants.EstuaryApiBaseUrl+"/content/by-cid/"+cid,
		nil,
	)
	var responseObject ByCID
	json.Unmarshal(response, &responseObject)
	log.Print("Data fetched from service: ", responseObject)
	return responseObject
}

func UploadContent(filePath string) UploadResponse {
	log.Print("Start upload data request")
	response := doMultipartApiRequest(
		"POST",
		constants.EstuaryApiShuttleBaseUrl+"/content/add",
		filePath,
	)
	var responseObject UploadResponse
	json.Unmarshal(response, &responseObject)
	log.Print("Data received from upload request: ", responseObject)
	return responseObject
}

func DownloadContent(cid string) string {
	filepath := "assets/downloaded.bin"

	// Create blank file
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	log.Print("Start download data request")
	resp, err := client.Get(constants.EstuaryDownloadApiUrl + "/" + cid)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Print("Download data received")

	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatalf("file write err: %v", err.Error())
	}

	defer file.Close()

	return filepath
}

func doMultipartApiRequest(method string, url string, filePath string) []byte {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, _ := os.Open(filePath)
	defer file.Close()
	part1, _ := writer.CreateFormFile("data", filepath.Base(filePath))
	_, _ = io.Copy(part1, file)
	err := writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("authorization", "Bearer "+constants.EstuaryApiToken)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Return api call response: ", string(responseData))
	return responseData
}

func doApiRequest(method string, url string, body io.Reader) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Panic(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", "Bearer "+constants.EstuaryApiToken)
	response, err := client.Do(req)
	if err != nil {
		log.Panic(err.Error())
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Return api call response: ", string(responseData))
	return responseData
}
