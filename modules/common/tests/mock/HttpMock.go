package mock

import (
	"bytes"
	"physiobot/modules/common/logger"
	"io/ioutil"
	"net/http"
)

type MockHTTPClient struct{}

func (m *MockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	method := request.Method
	url := request.URL.Path
	responseJson := ""
	statusCode := 500

	logger.Info("MOCK HTTP CALL: " + method + ": " + url)

	switch method {
	case "POST":
		switch url {
		case "/jws_edths6/EDTHS6Service":
			var responseXml string
			responseXml = "<env:Envelope xmlns:env=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"><env:Header/><env:Body><ftn:EdtLookupReply xmlns:ftn=\"http://ftn.fedex.com\"><java:HighestSeverity xmlns:java=\"java:com.fedex.nxgen.trdt.v3.ientities\">SUCCESS</java:HighestSeverity><java:TransactionDetail xmlns:java=\"java:com.fedex.nxgen.trdt.v3.ientities\" xsi:nil=\"true\"/><java:Version xmlns:java=\"java:com.fedex.nxgen.trdt.v3.ientities\"><java:ServiceId xsi:nil=\"true\"/><java:Major>4</java:Major><java:Intermediate>1</java:Intermediate><java:Minor>9</java:Minor></java:Version><java:InCoTerm xmlns:java=\"java:com.fedex.nxgen.trdt.v3.ientities\">CIF</java:InCoTerm><java:CommodityTaxes xmlns:java=\"java:com.fedex.nxgen.trdt.v3.ientities\"><java:HarmonizedCode>030211</java:HarmonizedCode><java:Taxes><java:TaxType>DUTY</java:TaxType><java:EffectiveDate>2020-05-22</java:EffectiveDate><java:Name>MFN</java:Name><java:TaxableValue><java:Currency>USD</java:Currency><java:Amount>158.8</java:Amount></java:TaxableValue><java:Description xsi:nil=\"true\"/><java:Formula>Amount = 0.15 * customsValue</java:Formula><java:Amount><java:Currency>USD</java:Currency><java:Amount>23.82</java:Amount></java:Amount></java:Taxes><java:Taxes><java:TaxType>GENERAL_SALES_TAX</java:TaxType><java:EffectiveDate xsi:nil=\"true\"/><java:Name>VAT</java:Name><java:TaxableValue><java:Currency>USD</java:Currency><java:Amount>188.71</java:Amount></java:TaxableValue><java:Description xsi:nil=\"true\"/><java:Formula>Amount = DPV * 0.16</java:Formula><java:Amount><java:Currency>USD</java:Currency><java:Amount>30.19</java:Amount></java:Amount></java:Taxes><java:Taxes><java:TaxType>CUSTOMS_SURCHARGES</java:TaxType><java:EffectiveDate xsi:nil=\"true\"/><java:Name>CPF</java:Name> <java:TaxableValue><java:Currency>USD</java:Currency><java:Amount>164.89</java:Amount></java:TaxableValue><java:Description xsi:nil=\"true\"/><java:Formula>Amount = 0.008 * quantity</java:Formula><java:Amount><java:Currency>USD</java:Currency><java:Amount>1.32</java:Amount></java:Amount> </java:Taxes> </java:CommodityTaxes></ftn:EdtLookupReply></env:Body></env:Envelope>"
			responseBody := ioutil.NopCloser(bytes.NewReader([]byte(responseXml)))
			return &http.Response{
				Status:     "OK",
				StatusCode: 200,
				Body:       responseBody,
			}, nil
		}
	}

	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(responseJson)))
	return &http.Response{
		StatusCode: statusCode,
		Body:       responseBody,
	}, nil
}
