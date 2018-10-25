package ethwatcher

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}

func HexToAddress(sHexAddr string) common.Address {
	return common.HexToAddress(sHexAddr)
}

func GetAddressFromPub(sPub string) (common.Address, []byte, error) {
	bPub, err := hex.DecodeString(sPub)
	if err != nil {
		return common.Address{}, bPub, err
	}

	pubKey, err := crypto.UnmarshalPubkey(bPub)
	if err != nil {
		return common.Address{}, bPub, err
	}

	return crypto.PubkeyToAddress(*pubKey), bPub, nil
}

func CreateRemoteSignOpts(sPub string, sPubHash string, sSignUrl string, sSrvId string, sSrvKey string) *bind.TransactOpts {
	keyAddr, bPub, _ := GetAddressFromPub(sPub)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, fmt.Errorf("not authorized to sign this account")
			}
			bTxHash := signer.Hash(tx).Bytes()
			postData, err := createPostData(hex.EncodeToString(bTxHash), sPubHash, sSrvId)
			if err != nil {
				return nil, err
			}
			postSig, err := SignPostData(postData, sSrvKey)
			if err != nil {
				return nil, err
			}
			sigData, err := RemoteSign(sSignUrl, postData, postSig, sSrvId)
			if err != nil {
				return nil, err
			}
			signature, err := CreateEthSignature(bPub, bTxHash, sigData)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

}
