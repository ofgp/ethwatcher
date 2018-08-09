package ethwatcher

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
)

func SignPostData(postData []byte, sSrvKey string) (string, error) {
	bSrvKey, err := hex.DecodeString(sSrvKey)
	if err != nil {
		ewLogger.Error("decode srvKey error", "error", err)
		return "", err
	}
	srvPrivKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), bSrvKey)
	dataHash := sha256.New()
	dataHash.Write(postData)
	sig, err := srvPrivKey.Sign(dataHash.Sum(nil))
	if err != nil {
		ewLogger.Error("sign postData error", "error", err)
		return "", err
	}
	return hex.EncodeToString(sig.Serialize()), nil
}
