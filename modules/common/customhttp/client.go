package customhttp

import (
	"crypto/tls"
	"net/http"
)

type HTTPClient interface {
	//Get(url string, headers map[string] string, response interface{}) (int, interface{}, error)
	//Post(url string, headers map[string] string, body interface{}) (int, []byte, error)
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	//Client = &http.Client{}
	Client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
