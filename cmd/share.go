package encloud

import (
	"encoding/json"
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"filecoin-encrypted-data-storage/types"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func ShareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "share",
		Short: "Share uploaded content to other user",
		Long:  `Share uploaded content with your cid and dek to another user`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _ := config.LoadConf("./config.yml")

			kek := ""
			privateKey := ""
			publicKey, _ := cmd.Flags().GetString("publicKey")
			pk, _ := cmd.Flags().GetString("privateKey")
			uuid, _ := cmd.Flags().GetString("uuid")
			email, _ := cmd.Flags().GetString("email")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("readPublicKeyFromPath")
			readPrivateKeyFromPath, _ := cmd.Flags().GetBool("readPrivateKeyFromPath")
			if readPublicKeyFromPath {
				kek = thirdparty.ReadKeyFile(publicKey)
			} else {
				kek = publicKey
			}

			if readPrivateKeyFromPath {
				privateKey = thirdparty.ReadKeyFile(pk)
			} else {
				privateKey = pk
			}

			fileMetaData := service.FetchByCid(kek + ":" + uuid)
			decryptedDek, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(decryptedDek)

			// Writing decryption dek
			err = ioutil.WriteFile("assets/dek.txt", decryptedDek, 0777)
			if err != nil {
				log.Fatalf("write file err: %v", err.Error())
			}

			subject := "Share content"
			r := service.NewRequest([]string{email}, subject, cfg)
			r.Send("./templates/share.html", map[string]string{"cid": fileMetaData.Cid[0]})

			response := types.RetrieveByCIDContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content shared successfully.",
				Data:       fileMetaData,
			}
			encoded, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	cmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	cmd.Flags().StringP("privateKey", "k", "", "Enter your private key")
	cmd.Flags().StringP("uuid", "u", "", "Enter your uuid")
	cmd.Flags().StringP("email", "e", "", "Enter email which you want to share")
	cmd.Flags().BoolP("readPublicKeyFromPath", "r", false, "Do you want public key read from path you have entered?")
	cmd.Flags().BoolP("readPrivateKeyFromPath", "o", false, "Do you want private key read from path you have entered?")
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("privateKey")
	cmd.MarkFlagRequired("uuid")
	cmd.MarkFlagRequired("email")
	return cmd
}

func init() {
	RootCmd.AddCommand(ShareCmd())
}
