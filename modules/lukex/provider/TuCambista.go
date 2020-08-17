package provider

import (
	"encoding/json"
	"net/http"
	//"physiobot/modules/common/logger"
	"strconv"
	"time"
)

type TuCambistaInterface Provider

type TuCambistaProvider struct {
	Name string
	PublicUrl string
	Url string
}

func (prov *TuCambistaProvider) GetName() string {
	prov.Name = "TuCambista"
	return prov.Name
}

func (prov *TuCambistaProvider) GetData() string {
	prov.Url = "https://app.tucambista.pe/"
	var timestamp = time.Now().Unix()
	var finalUrl = prov.Url + "api/transaction/getquote/500/USD/BUY/?_=" + strconv.FormatInt(timestamp, 10);
	//logger.Info(finalUrl)

	resp, err := http.Get(finalUrl)
	if err != nil {
		print(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	//logger.Info(result)

	defer resp.Body.Close()

	finalValue := 0.0
	if (result["id"] != "") {
		finalValue = result["exchangeRateUsed"].(float64)
	}

	return strconv.FormatFloat(finalValue, 'f', 6, 64)
}