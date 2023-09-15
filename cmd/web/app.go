package main

import (
	"context"
	"net/http"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/api"
	"github.com/encloud-tech/encloud/pkg/types"

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
	} else {
		response = types.GenerateKeyPairResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Keys generated successfully.",
			Data:       keys,
		}
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

func (a *App) SelectDirectory() string {
	file, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
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
	} else {
		response = types.UploadContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content uploaded successfully.",
			Data:       types.Uuid{Uuid: uuid},
		}
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

// Fetch data from db
func (a *App) ListKeys() types.ListKeysResponse {
	var response types.ListKeysResponse
	keys, err := api.ListKeys()
	if err != nil {
		response = types.ListKeysResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.ListKeys{},
		}
	} else {
		response = types.ListKeysResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content fetched successfully.",
			Data:       keys,
		}
	}

	return response
}

// Retrieve data by uuid
func (a *App) RetrieveByUUID(uuid string, kek string, privateKey string, retrievalFileStoragePath string) types.RetrieveByUUIDContentResponse {
	var response types.RetrieveByUUIDContentResponse
	fileMetaData, err := api.RetrieveByUUID(uuid, kek, privateKey, retrievalFileStoragePath)
	if err != nil {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileMetadata{},
		}
	} else {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content fetched successfully.",
			Data:       fileMetaData,
		}
	}

	return response
}

// Share data via email
func (a *App) Share(uuid string, kek string, privateKey string, email string) types.RetrieveByUUIDContentResponse {
	var response types.RetrieveByUUIDContentResponse
	fileMetaData, err := api.Share(uuid, kek, privateKey, email)
	if err != nil {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.FileMetadata{},
		}
	} else {
		response = types.RetrieveByUUIDContentResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content shared successfully.",
			Data:       fileMetaData,
		}
	}

	return response
}

// Retrieve shared content
func (a *App) RetrieveSharedContent(decryptedDekPath string, dekType string, cid string, fileName string, retrievalFileStoragePath string) types.SharedResponse {
	var response types.SharedResponse
	err := api.RetrieveSharedContent(decryptedDekPath, dekType, cid, fileName, retrievalFileStoragePath)
	if err != nil {
		response = types.SharedResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	} else {
		response = types.SharedResponse{
			Status:     "success",
			StatusCode: http.StatusCreated,
			Message:    "Content fetched successfully.",
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
	} else {
		response = types.ConfigResponse{
			Status:     "success",
			StatusCode: http.StatusFound,
			Message:    "Configuration saved successfully",
			Data:       types.ConfYaml{},
		}
	}

	return response
}

// Restore default config
func (a *App) RestoreDefaultConfig() types.ConfigResponse {
	var response types.ConfigResponse
	err := config.LoadDefaultConf()
	if err != nil {
		response = types.ConfigResponse{
			Status:     "fail",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       types.ConfYaml{},
		}
	} else {
		response = types.ConfigResponse{
			Status:     "success",
			StatusCode: http.StatusFound,
			Message:    "Configuration saved successfully",
			Data:       types.ConfYaml{},
		}
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
	} else {
		response = types.ConfigResponse{
			Status:     "success",
			StatusCode: http.StatusFound,
			Message:    "Config data fetched successfully",
			Data:       conf,
		}
	}

	return response
}
