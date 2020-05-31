package customhttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"physiobot/modules/common/logger"
	"io/ioutil"
	"net/http"
)

func Get(url string, headers map[string]string) (int, []byte, error) {
	statusCode := 500
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return statusCode, nil, err
	}

	for key, element := range headers {
		request.Header.Set(key, element)
	}

	result, err := Client.Do(request)
	if result != nil {
		statusCode = result.StatusCode
	}
	if err != nil {
		return statusCode, nil, err
	}

	defer result.Body.Close()
	resultData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return statusCode, nil, err
	}

	return statusCode, resultData, nil
}

func Post(url string, headers map[string]string, body interface{}, format string) (int, []byte, error) {
	statusCode := 500

	request, err := BuildRequest("POST", url, headers, body, format)

	if err != nil {
		return statusCode, nil, err
	}

	response, err := Client.Do(request)
	if response != nil {
		statusCode = response.StatusCode
	}
	if err != nil {
		return statusCode, nil, err
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return statusCode, nil, err
	}

	return statusCode, responseData, nil
}

func Delete(url string, headers map[string]string, body interface{}) (int, []byte, error) {

	logger.Info(url)
	statusCode := 500

	request, err := http.NewRequest("DELETE", url, nil)
	/*
		buf := new(bytes.Buffer)
		if body != nil {
			json.NewEncoder(buf).Encode(body)
			request, err := http.NewRequest("DELETE", url, buf)
		}
	*/

	if err != nil {
		return statusCode, nil, err
	}

	for key, element := range headers {
		request.Header.Set(key, element)
	}

	response, err := Client.Do(request)
	logger.Info(response)
	logger.Info(response.StatusCode)
	if response != nil {
		statusCode = response.StatusCode
	}
	if err != nil {
		return statusCode, nil, err
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return statusCode, nil, err
	}

	return statusCode, responseData, nil
}

func BuildRequest(method string, url string, headers map[string]string, body interface{}, format string) (*http.Request, error) {
	var err error
	var request *http.Request

	switch format {
	case "json":
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return request, err
		}
		request, err = http.NewRequest(method, url, buf)
	case "xml":
		out, err := xml.MarshalIndent(body, "\t", "\t")
		if err != nil {
			return request, err
		}
		request, err = http.NewRequest(method, url, bytes.NewReader(out))
	}

	if err != nil {
		return request, err
	}

	for key, element := range headers {
		request.Header.Set(key, element)
	}

	return request, nil
}