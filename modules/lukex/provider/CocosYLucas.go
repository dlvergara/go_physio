package provider

import (
	"encoding/json"
	"net/http"
	"net/url"
	//"physiobot/modules/common/logger"
	"strings"
)

type CocosYLucasInterface Provider

type CocosYLucasProvider struct {
	Name string
	PublicUrl string
	Url string
}

func (prov *CocosYLucasProvider) GetName() string {
	prov.Name = "CocosyLucas"
	return prov.Name
}

func (prov *CocosYLucasProvider) GetToken() string {
	client := &http.Client{}

	postUrl := "https://www.cocosylucasbcp.com/toc"
	reqBody := url.Values{
		"fn": {"obtenerToken"},
	}
	req, err := http.NewRequest("POST", postUrl , strings.NewReader(reqBody.Encode()))
	if err != nil {
		print(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("charset", "utf-8")
	req.Header.Add("app-code", "MY")
	req.Header.Add("referer", "https://www.cocosylucasbcp.com/")
	req.Header.Add("origin", "https://www.cocosylucasbcp.com/")
	req.Header.Add("dnt", "1")

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	//logger.Info(result)

	defer resp.Body.Close()

	return result["access_token"].(string)
}

func (prov *CocosYLucasProvider) GetData() string {
	client := &http.Client{}
	prov.Url = "https://www.cocosylucasbcp.com/poly/currency-exchanges"

	token := prov.GetToken()
	//logger.Info(token)

	req, err := http.NewRequest("GET", prov.Url, nil)
	if err != nil {
		print(err)
	}
	req.Header.Add("authorization", "Bearer " + token)
	//req.Header.Add("charset", "utf-8")
	req.Header.Add("app-code", "MY")
	req.Header.Add("referer", "https://www.cocosylucasbcp.com/")
	req.Header.Add("origin", "https://www.cocosylucasbcp.com/")
	req.Header.Add("dnt", "1")

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	//logger.Info(result)

	defer resp.Body.Close()

	//logger.Info(resp.StatusCode)

	var finalValue = ""
	if(resp.StatusCode == 200) {
		if (result["currencyExchangeList"] != nil) {
			//var exchanges = result["currencyExchangeList"].(map[string]string);

			/*
			for _, element := range exchanges {
				if(element["maxUsdPurchase"].(string) == "2499.9999") {

				}
			}

			for (final i in exchanges) {
				if (i['maxUsdPurchase'] == "2499.9999") {
					result = i['rateSale'];
					break;
				}
			}
			*/
		} else {
			var status = result["status"].(map[string]interface{})
			finalValue = status["description"].(string)
		}
	}

	//return strconv.FormatFloat(finalValue, 'f', 6, 64)
	return finalValue
}