package cs_sdk

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RequestRaw(path string, method string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", URL, VERSION, path)

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"api_key":      {os.Getenv("CS_API_KEY")},
		"access_token": {os.Getenv("CS_DELIVERY_TOKEN")},
		"branch":       {os.Getenv("CS_BRANCH")},
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Request(path string, method string) (string, error) {
	res, err := RequestRaw(path, method)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	return string(body), nil
}
