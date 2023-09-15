package api

import (
	"os"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"
)

func GenerateKeyPair(kekType string) (types.Keys, error) {
	var keys types.Keys

	if kekType == "rsa" {
		os.RemoveAll(config.DotKeys)
		thirdparty.InitCrypto()
		keys = types.Keys{PublicKey: thirdparty.GetIdRsaPubStr(), PrivateKey: thirdparty.GetIdRsaStr()}
	} else if kekType == "ecies" {
		k := thirdparty.EciesGenerateKeyPair()
		keys = types.Keys{PublicKey: k.PublicKey.Hex(false), PrivateKey: k.Hex()}
	} else {
		return types.Keys{}, nil
	}

	return keys, nil

}
