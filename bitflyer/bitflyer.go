package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://bitflyer.com/v1/"

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	apiClient := &APIClient{key, secret, &http.Client{}}
	return apiClient
}

func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println(timestamp)
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	// baseURLのparse
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	// urlPathのparse
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	// baseURLとurlPathからendpointを生成
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	// http requestを作成
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	// qにクエリを設定
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	// http requestのRawQueryに設定したクエリを設定
	req.URL.RawQuery = q.Encode()
	// http headerを設定
	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}
	// http requestを送信
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// 最後にhttp response Bodyはクローズ
	defer resp.Body.Close()
	// http response Bodyを読み込む
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Amount      float64 `json:"amount"`
	Available   float64 `json:"available"`
}

func (api *APIClient) GetBalance() ([]Balance, error) {
	url := "me/getbalance"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}
