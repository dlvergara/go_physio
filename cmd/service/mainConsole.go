package main

import (
	"physiobot/config"
	"physiobot/modules/common/db"
	"physiobot/modules/common/logger"
	"physiobot/modules/lukex/model"
	"physiobot/modules/lukex/provider"
	"strconv"
	"time"
)

func main() {
	dt := time.Now()
	logger.Info("Start process " + dt.String())
	configuration := config.GetConfig()
	logger.Info(configuration)

	providers := make(map[int]provider.Provider)
	responses := make(map[int]string)
	//providers[1] = new(provider.TuCambistaProvider)
	providers[2] = new(provider.CocosYLucasProvider)

	for index, providerObj := range providers {
		logger.Info(providerObj.GetName())
		var res = providerObj.GetData()
		responses[index] = res

		if(res != "") {

			exchangeModel := new(model.ExchangeModel)
			exchangeModel.Provider = "lukex_" + providerObj.GetName()
			exchangeModel.Timestamp = time.Now().Format("2006-01-02 03:04:05")

			_, err := strconv.ParseFloat(res, 64)
			if (err != nil) {
				logger.Error(err)
				exchangeModel.Exchange = "20"
				exchangeModel.Note = res
			} else {
				exchangeModel.Exchange = res
				exchangeModel.Note = ""
			}

			// Create
			db.DbLukex().Table("exchange").Create(exchangeModel)
		} else {
			logger.Error(providerObj.GetName() + " - empty response")
		}
	}

	logger.Info(responses)
}