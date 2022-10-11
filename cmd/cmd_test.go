package cmd

import (
	"bytes"
	"encoding/json"
	"filecoin-encrypted-data-storage/types"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPairCommand(t *testing.T) {
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
	filePath := "./../testdata/Provider2InputData.csv"
	uploadContentCmd := UploadContentCmd()
	uploadContentBuf := new(bytes.Buffer)
	uploadContentCmd.SetOut(uploadContentBuf)
	uploadContentCmd.SetErr(uploadContentBuf)
	uploadContentCmd.SetArgs([]string{"-p", publicKey, "-f", filePath})
	uploadContentCmd.Execute()
	var uploadContentResponseObject types.UploadContentResponse
	json.Unmarshal(uploadContentBuf.Bytes(), &uploadContentResponseObject)
	assert.NotNil(t, uploadContentResponseObject.Data)
	CID := uploadContentResponseObject.Data.CID
	log.Println("cid: " + CID)

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

	// Finally, we have retrieved uploaded content using cid.
	retrieveContentByCidBuf := new(bytes.Buffer)
	retrieveContentByCidCmd := RetrieveByCidCmd()
	retrieveContentByCidCmd.SetOut(retrieveContentByCidBuf)
	retrieveContentByCidCmd.SetErr(retrieveContentByCidBuf)
	retrieveContentByCidCmd.SetArgs([]string{"-p", publicKey, "-k", privateKey, "-c", CID})
	retrieveContentByCidCmd.Execute()
	var retrieveContentByCidResponseObject types.RetrieveByCIDContentResponse
	json.Unmarshal(retrieveContentByCidBuf.Bytes(), &retrieveContentByCidResponseObject)
	assert.NotNil(t, retrieveContentByCidResponseObject.Data)
	log.Println(retrieveContentByCidResponseObject.Data)
}
