package okx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/wsy0414/go-okx/internal/util"
)

type OkxClient struct {
	AccessKey        string
	SecretKey        string
	AccessPasspharse string
}

func NewOkxClient(accessKey, secretKey, accessPasspharse string) *OkxClient {
	return &OkxClient{
		AccessKey:        accessKey,
		SecretKey:        secretKey,
		AccessPasspharse: accessPasspharse,
	}
}

func (client *OkxClient) get(path string, params *url.Values, resBody any) error {
	u, _ := url.ParseRequestURI(domain)
	u.Path = path
	if params != nil {
		u.RawQuery = params.Encode()
	}

	if err := client.sendRequest(http.MethodGet, u, nil, resBody); err != nil {
		return err
	}

	return nil
}

func (client *OkxClient) sendRequest(method string, url *url.URL, reqBody []byte, resBody any) error {
	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	client.setDefaultHeader(&req.Header, method, strings.Replace(url.String(), "domain", "", -1), reqBody)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	data, err := readBody(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = json.Unmarshal(data, &resBody); err != nil {
		return err
	}

	return err
}

func readBody(body io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *OkxClient) setDefaultHeader(header *http.Header, method string, url string, body any) {
	now := util.GetISOFormat(time.Now())
	header.Set(OKX_HEADER_ACCESS_KEY, client.AccessKey)
	signStr, err := client.sign(body, now, method, url)
	if err != nil {
		panic(err)
	}
	header.Set(OKX_HEADER_ACCESS_SIGN, signStr)
	header.Set(OKX_HEADER_ACCESS_TIMESTAMP, now)
	header.Set(OKX_HEADER_ACCESS_PASSPHRASE, client.AccessPasspharse)
	if method == http.MethodPost {
		header.Set("Content-Type", "application/json")
	}
}

func (client *OkxClient) sign(body any, timeString string, methodStr string, url string) (string, error) {
	bodyStr := ""
	if body != nil {
		jsonByte, err := json.Marshal(body)
		if err != nil {
			return "", err
		}

		bodyStr = string(jsonByte)
	}

	s := timeString + methodStr + url + bodyStr
	b, err := util.ToSha256(s, client.SecretKey)
	if err != nil {
		return "", err
	}

	return util.ToBase64(b), nil
}
