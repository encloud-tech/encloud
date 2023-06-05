package api

import (
	"encloud/config"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"os"
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
