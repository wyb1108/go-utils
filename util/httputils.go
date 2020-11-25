package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpPostJson(url string, data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return respBytes, nil
}
