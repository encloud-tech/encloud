package service

import (
	"bytes"
	"encoding/json"
	"filecoin-encrypted-data-storage/config"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"filecoin-encrypted-data-storage/types"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

func (e *Estuary) UploadContent(filePath string) types.UploadResponse {
	log.Print("Start upload data request")
	response := e.doMultipartApiRequest(
		"POST",
		e.config.Estuary.ShuttleApiUrl+"/content/add",
		filePath,
	)
	var responseObject types.UploadResponse
	json.Unmarshal(response, &responseObject)
	log.Print("Data received from upload request: ", responseObject)
	return responseObject
}

func (e *Estuary) ChunkUploadContent(filePath string, chunkSize int64) []string {
	log.Print("Start upload data request")
	var cids []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		rand.Seed(time.Now().UnixNano())
		var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		rs := thirdparty.RandomString(20, alphabet)

		err = ioutil.WriteFile("assets/"+rs+".csv", buffer[:bytesread], 0777)
		if err != nil {
			log.Fatalf("write file err: %v", err.Error())
			// return err
		}

		response := e.doMultipartApiRequest(
			"POST",
			e.config.Estuary.ShuttleApiUrl+"/content/add",
			filePath,
		)
		var responseObject types.UploadResponse
		json.Unmarshal(response, &responseObject)
		log.Print("Data received from upload request: ", responseObject)
		if responseObject.CID != "" {
			os.Remove("assets/" + rs + ".csv")
			cids = append(cids, "https://dweb.link/ipfs/"+responseObject.CID)
		}
	}

	return cids
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

func (e *Estuary) doMultipartApiRequest(method string, url string, filePath string) []byte {
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
	req.Header.Set("authorization", "Bearer "+e.config.Estuary.Token)
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
