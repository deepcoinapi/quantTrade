package dc

import (
	//"apiRequest/structs"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Sign struct {
	Url        string
	Key        string
	SecretKey  string
	Passphrase string
}

func NewSign() *Sign {
	s := &Sign{
		Url: URL,
	}
	if v := os.Getenv("KEY"); v != "" {
		s.Key = v
	}
	if v := os.Getenv("SECRETKEY"); v != "" {
		s.SecretKey = v
	}
	if v := os.Getenv("PASSPHRASS"); v != "" {
		s.Passphrase = v
	}

	return s
}

func DoSign(timestamp string, method string, requestPath string, body string, secretKey string) (sign string, err error) {
	message := fmt.Sprintf("%s%s%s", timestamp, method, requestPath)
	if body != "" {
		message = fmt.Sprintf(message+"%s", body)
	}

	//fmt.Println(fmt.Sprintf("明文: %s", message))

	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(message))
	digest := hash.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString(digest)

	//fmt.Println(fmt.Sprintf("密文: %s", encoded))

	return encoded, nil
}

func DoHttp(requestURL string, requestMethod string, requestPath string, requestBody string, env *Sign) (body []byte, err error) {
	var req *http.Request
	//var err error
	if requestMethod == "POST" {
		reqBody := bytes.NewBufferString(requestBody)
		req, err = http.NewRequest(requestMethod, requestURL, reqBody)
	} else {
		req, err = http.NewRequest(requestMethod, requestURL, nil)
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("http new request error:%v", err))
		return nil, err
	}

	if env != nil {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		sign, err := DoSign(timestamp, requestMethod, requestPath, requestBody, env.SecretKey)
		if err != nil {
			fmt.Println(fmt.Sprintf("sign error:%v", err))
			return nil, err
		}

		req.Header.Set("DC-ACCESS-KEY", env.Key)
		req.Header.Set("DC-ACCESS-SIGN", sign)
		req.Header.Set("DC-ACCESS-TIMESTAMP", timestamp)
		req.Header.Set("DC-ACCESS-PASSPHRASE", env.Passphrase)
	}

	//fmt.Println(fmt.Sprintf("request: %v", req))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Sprintf("http request error:%v", err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("read response error:%v", err))
		return nil, err
	}

	fmt.Println(string(body))
	return body, nil
}
