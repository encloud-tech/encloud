package api

import (
	"encloud/config"
	"encloud/pkg/service"
	thirdparty "encloud/third_party"
	"os"
)

func RetrieveSharedContent(decryptedDekPath string, dekType string, cid string) error {
	cfg, err := config.LoadConf("./config.yaml")
	if err != nil {
		// Load default configuration from config.go file if config.yaml file not found
		cfg, _ = config.LoadConf()
	}
	estuaryService := service.New(cfg)

	dek := thirdparty.ReadFile(decryptedDekPath)

	filepath := estuaryService.DownloadContent(cid)
	if dekType == "aes" {
		err := thirdparty.DecryptWithAES(dek, filepath, "assets/decrypted.csv")
		if err != nil {
			return err
		}
	} else {
		err := thirdparty.DecryptWithChacha20poly1305(dek, filepath, "assets/decrypted.csv")
		if err != nil {
			return err
		}
	}

	os.Remove("assets/downloaded.bin")
	return nil
}
