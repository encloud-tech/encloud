package main

import (
	"context"
	"encloud/pkg/api"
	"encloud/pkg/types"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Generate key pair
func (a *App) GenerateKeyPair(kekType string) types.GenerateKeyPairResponse {
	var keys types.Keys
	var response types.GenerateKeyPairResponse
	keys, err := api.GenerateKeyPair(kekType)
	if err != nil {
		response = types.GenerateKeyPairResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.Keys{},
		}
	}

	response = types.GenerateKeyPairResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Keys generated successfully.",
		Data:       keys,
	}

	return response
}

func (a *App) SelectFile() string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return err.Error()
	}
	return file
}

// Upload data to esatury
func (a *App) Upload(filePath string, kekType string, dekType string, kek string) types.UploadContentResponse {
	var response types.UploadContentResponse
	uuid, err := api.Upload(filePath, kekType, dekType, kek)
	if err != nil {
		response = types.UploadContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.Uuid{},
		}
	}

	response = types.UploadContentResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Content uploaded successfully.",
		Data:       types.Uuid{Uuid: uuid},
	}

	return response
}

// Fetch data from db
func (a *App) List(kek string) types.ListContentResponse {
	var response types.ListContentResponse
	fileData := api.List(kek)
	response = types.ListContentResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Content fetched successfully.",
		Data:       fileData,
	}

	return response
}

// Retrieve data by uuid
func (a *App) RetrieveByUUID(uuid string, kek string, privateKey string) types.RetrieveByCIDContentResponse {
	var response types.RetrieveByCIDContentResponse
	fileMetaData, err := api.RetrieveByUUID(uuid, kek, privateKey)
	if err != nil {
		response = types.RetrieveByCIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileMetadata{},
		}
	}
	response = types.RetrieveByCIDContentResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Content fetched successfully.",
		Data:       fileMetaData,
	}

	return response
}

// Share data via email
func (a *App) Share(uuid string, kek string, privateKey string, email string) types.RetrieveByCIDContentResponse {
	var response types.RetrieveByCIDContentResponse
	fileMetaData, err := api.Share(uuid, kek, privateKey, email)
	if err != nil {
		response = types.RetrieveByCIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileMetadata{},
		}
	}
	response = types.RetrieveByCIDContentResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Content shared successfully.",
		Data:       fileMetaData,
	}

	return response
}

// Retrieve shared content
func (a *App) RetrieveSharedContent(decryptedDekPath string, dekType string, cid string) types.ErrorResponse {
	var response types.ErrorResponse
	err := api.RetrieveSharedContent(decryptedDekPath, dekType, cid)
	if err != nil {
		response = types.ErrorResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response
}

// Store config
func (a *App) StoreConfig(conf types.ConfYaml) types.ConfigResponse {
	var response types.ConfigResponse
	err := api.Store(conf)
	if err != nil {
		response = types.ConfigResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.ConfYaml{},
		}
	}

	response = types.ConfigResponse{
		Status:     "success",
		StatusCode: http.StatusFound,
		Message:    "Configuration saved successfully",
		Data:       types.ConfYaml{},
	}

	return response
}

// Fetch config
func (a *App) FetchConfig() types.ConfigResponse {
	var response types.ConfigResponse
	conf, err := api.Fetch()
	if err != nil {
		response = types.ConfigResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.ConfYaml{},
		}
	}

	response = types.ConfigResponse{
		Status:     "success",
		StatusCode: http.StatusFound,
		Message:    "Config data fetched successfully",
		Data:       conf,
	}

	return response
}