package net

import(
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"bytes"
	"io"
	"errors"
	"fmt"
)

func HttpPostRequest(apiUrl string, bytesData []byte, headers map[string]string, httpClient *http.Client)  (*http.Response, []byte, error){
	var err error

	var isGzip = false
	if len(headers) > 0 {
		for key, item := range headers {
			if key == "Accept-Encoding" && item == "gzip" {
				isGzip = true
				break
			}
		}
	}

	var reader *bytes.Buffer

	if isGzip {
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if _, err = zw.Write(bytesData); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("gzip error='%s'", err))
		}
		zw.Close()
		reader = &zBuf
	} else {
		reader = bytes.NewBuffer(bytesData)

	}

	request, err := http.NewRequest("POST", apiUrl, reader)

	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("newRequest error='%s'", err))
	}

	if len(headers) > 0 {
		for key, item := range headers {
			request.Header.Set(key, item)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()

	body := response.Body

	if response.Header.Get("Content-Encoding") == "gzip" {
		body2, e := gzip.NewReader(response.Body)
		if e == nil {
			body = body2
		} else if e != io.EOF{
			return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", e))
		}
	}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}

func HttpGetRequest(apiUrl string, httpClient *http.Client)  (*http.Response, []byte, error) {
	response, err := httpClient.Get(apiUrl)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, body, nil
}