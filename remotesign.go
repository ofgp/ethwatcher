package ethwatcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignInput struct {
	InputHex   string `json:"inputHex"`
	PubHashHex string `json:"pubHashHex"`
	ServiceId  string `json:"serviceId"`
	TimeStamp  int64  `json:"timestamp"`
}

//SignData 签名结果
type SignData struct {
	R               string `json:"r"`
	S               string `json:"s"`
	SignatureDerHex string `json:"signatureDerHex"`
}

//SignResponse 秘钥服务签名返回结构
type SignResponse struct {
	Code int       `json:"code"`
	Data *SignData `json:"data"`
	Msg  string    `json:"msg"`
}

func createPostData(sTxHash string, sPubHash string, sSrvId string) ([]byte, error) {
	inputData := SignInput{
		InputHex:   sTxHash,
		PubHashHex: sPubHash,
		ServiceId:  sSrvId,
		TimeStamp:  time.Now().Unix(),
	}

	postData, err := json.Marshal(&inputData)
	if err != nil {
		ewLogger.Error("marshal postData error", "error", err)
		return nil, err
	}
	return postData, nil
}

//Sign 使用签名服务对待签名数据进行签名
func RemoteSign(sSignUrl string, postData []byte, sSignature string, sSrvId string) (*SignData, error) {
	client := &http.Client{}

	signUrl := strings.Join([]string{sSignUrl, "key", "sign"}, "/")
	req, err := http.NewRequest("POST", signUrl, bytes.NewBuffer(postData))
	if err != nil {
		ewLogger.Error("create Request error", "error", err)
		return nil, err
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("signature", sSignature)
	req.Header.Set("serviceId", sSrvId)

	resp, err := client.Do(req)

	if err != nil {
		ewLogger.Error("do request error", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ewLogger.Error("read response body error", "error", err)
		return nil, err
	}

	var sigRes SignResponse
	err = json.Unmarshal(body, &sigRes)
	if err != nil {
		ewLogger.Error("unmarshal response body error", "error", err)
		return nil, err
	}

	if sigRes.Code == 0 {
		return sigRes.Data, nil
	}

	ewLogger.Error(fmt.Sprintf("response code error, code is %v", sigRes.Code))
	return nil, fmt.Errorf("error response code is %d", sigRes.Code)
}

func CreateEthSignature(bPub []byte, bTxHash []byte, sigData *SignData) ([]byte, error) {
	rb := common.LeftPadBytes(common.FromHex(sigData.R), 32)
	sb := common.LeftPadBytes(common.FromHex(sigData.S), 32)

	length := 1 + len(rb) + len(sb)
	b := make([]byte, length)

	offsetR := copy(b[0:], rb)
	offsetS := copy(b[offsetR:], sb)
	for recId := 0; recId < 4; recId++ {
		b[offsetR+offsetS] = byte(recId)
		recoverPub, err := crypto.Ecrecover(bTxHash, b)
		if err == nil && bytes.Equal(bPub, recoverPub) {
			return b, nil
		}
	}
	ewLogger.Error("recover id is missing")
	return b, fmt.Errorf("[ethwatcher CreateEthSignature] recover id is missing")
}
