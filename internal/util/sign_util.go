package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ToSha256(str string, secretKey string) ([]byte, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(str))
	if err != nil {
		return nil, err
	}

	return mac.Sum(nil), nil
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
