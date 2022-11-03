package service

import (
	"bytes"
	"encloud/config"
	"encloud/types"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// New func implements the storage interface
func New(config *config.ConfYaml) *Estuary {
	return &Estuary{
		config: config,
	}
}

type Estuary struct {
	config *config.ConfYaml
}

func (e *Estuary) FetchAllContents() types.Contents {
	log.Print("Start fetching data from service")
	response := e.doApiRequest(
		"GET",
		e.config.Estuary.BaseApiUrl+"/content/list",
		nil,
	)
	var responseObject types.Contents
	json.Unmarshal(response, &responseObject)
	log.Print("Data fetched from service: ", responseObject)
	return responseObject
}

func (e *Estuary) FetchContentByCid(cid string) types.ByCID {
	log.Print("Start fetching data from service")
	response := e.doApiRequest(
		"GET",
		e.config.Estuary.BaseApiUrl+"/content/by-cid/"+cid,
		nil,
	)
	var responseObject types.ByCID
	json.Unmarshal(response, &responseObject)
	log.Print("Data fetched from service: ", responseObject)
	return responseObject
}

func (e *Estuary) UploadContent(filePath string) (types.UploadResponse, error) {
	log.Print("Start upload data request")
	var responseObject types.UploadResponse
	response, err := e.doMultipartApiRequest(
		"POST",
		e.config.Estuary.ShuttleApiUrl+"/content/add",
		filePath,
	)
	if err != nil {
		return responseObject, err
	}
	json.Unmarshal(response, &responseObject)
	log.Print("Data received from upload request: ", responseObject)
	return responseObject, nil
}

func (e *Estuary) DownloadContent(cid string) string {
	filepath := "assets/downloaded.bin"

	// Create blank file
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	log.Print("Start download data request")
	resp, err := client.Get(e.config.Estuary.DownloadApiUrl + "/" + cid)
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

func (e *Estuary) doMultipartApiRequest(method string, url string, filePath string) ([]byte, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, _ := os.Open(filePath)
	defer file.Close()
	part1, _ := writer.CreateFormFile("data", filepath.Base(filePath))
	_, _ = io.Copy(part1, file)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("authorization", "Bearer "+e.config.Estuary.Token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Print("Return api call response: ", string(responseData))
	return responseData, nil
}

func (e *Estuary) doApiRequest(method string, url string, body io.Reader) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Panic(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", "Bearer "+e.config.Estuary.Token)
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
