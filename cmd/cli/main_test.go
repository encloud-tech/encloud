package main

import (
	"bytes"
	"encloud/config"
	"encloud/pkg/types"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLICommands(t *testing.T) {
	if _, err := os.Stat(config.TestDir); err != nil {
		// create it
		err = os.Mkdir(config.TestDir, 0700)
		if err != nil {
			log.Fatalf("ERROR: fail to create test data dir, %s", err.Error())
			os.Exit(1)
		}
	}

	// Update default configuration
	configFilePath := "../../testdata/config.yaml"
	updateConfigBuf := new(bytes.Buffer)
	configCmd := ConfigCmd()
	configCmd.SetOut(updateConfigBuf)
	configCmd.SetErr(updateConfigBuf)
	configCmd.SetArgs([]string{"-f", configFilePath})
	configCmd.Execute()
	var configResponseObject types.ConfigResponse
	json.Unmarshal(updateConfigBuf.Bytes(), &configResponseObject)
	assert.NotNil(t, configResponseObject.Data)

	// First we have generate key pair to encrypt and decrypt dek.
	generateKeyPairBuf := new(bytes.Buffer)
	generateKeyPairCmd := GenerateKeyPairCmd()
	generateKeyPairCmd.SetOut(generateKeyPairBuf)
	generateKeyPairCmd.SetErr(generateKeyPairBuf)
	generateKeyPairCmd.Execute()
	var generateKeyPairResponseObject types.GenerateKeyPairResponse
	json.Unmarshal(generateKeyPairBuf.Bytes(), &generateKeyPairResponseObject)
	assert.NotNil(t, generateKeyPairResponseObject.Data)
	publicKey := generateKeyPairResponseObject.Data.PublicKey
	privateKey := generateKeyPairResponseObject.Data.PrivateKey
	log.Println("public key:" + publicKey)
	log.Println("private key:" + privateKey)

	// After that we can upload encrypted file using dek which is also encrypted using generated public key.
	filePath := "../../testdata/Provider2InputData.csv"
	uploadContentCmd := UploadContentCmd()
	uploadContentBuf := new(bytes.Buffer)
	uploadContentCmd.SetOut(uploadContentBuf)
	uploadContentCmd.SetErr(uploadContentBuf)
	uploadContentCmd.SetArgs([]string{"-p", publicKey, "-f", filePath, "-t", "aes"})
	uploadContentCmd.Execute()
	var uploadContentResponseObject types.UploadContentResponse
	json.Unmarshal(uploadContentBuf.Bytes(), &uploadContentResponseObject)
	assert.NotNil(t, uploadContentResponseObject.Data)
	Uuid := uploadContentResponseObject.Data.Uuid
	log.Println("Uuid: " + Uuid)

	// Now we fetch list of file meta data from database.
	listContentBuf := new(bytes.Buffer)
	listContentsCmd := ListContentsCmd()
	listContentsCmd.SetOut(listContentBuf)
	listContentsCmd.SetErr(listContentBuf)
	listContentsCmd.SetArgs([]string{"-p", publicKey})
	listContentsCmd.Execute()
	var listContentResponseObject types.ListContentResponse
	json.Unmarshal(listContentBuf.Bytes(), &listContentResponseObject)
	assert.NotNil(t, listContentResponseObject.Data)
	log.Println(listContentResponseObject.Data)

	// Finally, we have retrieved uploaded content using UUID.
	retrieveContentByUUIDBuf := new(bytes.Buffer)
	retrieveContentByUUIDCmd := RetrieveByUUIDCmd()
	retrieveContentByUUIDCmd.SetOut(retrieveContentByUUIDBuf)
	retrieveContentByUUIDCmd.SetErr(retrieveContentByUUIDBuf)
	retrieveContentByUUIDCmd.SetArgs([]string{"-p", publicKey, "-k", privateKey, "-u", Uuid, "-s", config.TestDir})
	retrieveContentByUUIDCmd.Execute()
	var retrieveContentByUUIDResponseObject types.RetrieveByUUIDContentResponse
	json.Unmarshal(retrieveContentByUUIDBuf.Bytes(), &retrieveContentByUUIDResponseObject)
	assert.NotNil(t, retrieveContentByUUIDResponseObject.Data)
	log.Println(retrieveContentByUUIDResponseObject.Data)

	// Share content via email.
	shareBuf := new(bytes.Buffer)
	shareCmd := ShareCmd()
	shareCmd.SetOut(shareBuf)
	shareCmd.SetErr(shareBuf)
	shareCmd.SetArgs([]string{"-p", publicKey, "-k", privateKey, "-u", Uuid, "-e", "test@encloud.test"})
	shareCmd.Execute()
	var shareResponseObject types.RetrieveByUUIDContentResponse
	json.Unmarshal(shareBuf.Bytes(), &shareResponseObject)
	assert.NotNil(t, shareResponseObject.Data)
	cid := shareResponseObject.Data.Cid[0]
	log.Println(shareResponseObject.Data)

	// Retrieve shared content.
	retrieveSharedContentBuf := new(bytes.Buffer)
	retrieveSharedContentCmd := RetrieveSharedContentCmd()
	retrieveSharedContentCmd.SetOut(retrieveSharedContentBuf)
	retrieveSharedContentCmd.SetErr(retrieveSharedContentBuf)
	retrieveSharedContentCmd.SetArgs([]string{"-c", cid, "-d", config.Assets + "/" + fmt.Sprint(shareResponseObject.Data.Timestamp) + "_dek.txt", "-s", config.TestDir, "-n", "shared.csv"})
	retrieveSharedContentCmd.Execute()
	var retrieveSharedContentResponseObject types.RetrieveByUUIDContentResponse
	json.Unmarshal(retrieveSharedContentBuf.Bytes(), &retrieveSharedContentResponseObject)
	assert.NotNil(t, retrieveSharedContentResponseObject.Data)
	log.Println(retrieveSharedContentResponseObject.Data)

	defer os.RemoveAll(config.TestDir)
}
