package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/optiopay/kafka/v2"
	"github.com/optiopay/kafka/v2/proto"
	"net/http"
	"physiobot/config"
	"physiobot/modules/common/db"
	"physiobot/modules/common/logger"
	"physiobot/modules/lukex/model"
	"physiobot/modules/lukex/rest"
	"strconv"
)

func ProcessData(c echo.Context) error {
	var kafkaAddrs = []string{"107.170.208.14:9092"}

	topics := map[int]string{}

	//topics[1] = "CocosyLucas"
	//topics[2] = "JetPeru"
	//topics[3] = "TKambio"
	//topics[4] = "CambistaInca"
	topics[1] = "TuCambista"

	configuration := config.GetConfig()
	fmt.Println(configuration)

	for _, element := range topics {
		//logger.Info(element)

		conf := kafka.NewBrokerConf("test-client")
		conf.AllowTopicCreation = false

		// connect to kafka cluster
		broker, err := kafka.Dial(kafkaAddrs, conf)
		if err != nil {
			logger.Error("cannot connect to kafka cluster: " + err.Error())
		}
		defer broker.Close()
		go printConsumed(broker, "lukex_" + element)
	}

	return c.JSON(http.StatusNoContent, "")
}

// printConsumed read messages from kafka and print them out
func printConsumed(broker kafka.Client, topic string) {
	logger.Info("Start consumer for: " + topic)

	conf := kafka.NewConsumerConf(topic, 0)
	conf.StartOffset = kafka.StartOffsetOldest

	consumer, err := broker.Consumer(conf)
	if err != nil {
		logger.Error("cannot create kafka consumer for " + topic + "--" + err.Error())
	} else {
		for {
			msg, err := consumer.Consume()
			if err != nil {
				if err != kafka.ErrNoData {
					logger.Info("cannot consume " + topic + " topic message: " + err.Error())
				}
				break
			}
			//logger.Info("message %f - $s", msg.Offset, msg.Value)//+ ": " +
			saveReceivedData(msg, topic)
		}
		logger.Info("consumer quit")
	}
}

func saveReceivedData(msg *proto.Message, topic string) {
	var offset = strconv.FormatInt(msg.Offset, 10)
	var jsonData = string(msg.Value[:])
	logger.Info("Mensaje " + offset + ": " + jsonData)

	bytes := []byte(jsonData)
	message := new(rest.Message)
	json.Unmarshal(bytes, &message)

	exchangeModel := new(model.ExchangeModel)
	exchangeModel.Provider = topic
	exchangeModel.Timestamp = message.Date
	exchangeModel.Exchange = message.Value

	//logger.Info(exchangeModel)
	// Create
	db.DbLukex().Table("exchange").Create(exchangeModel)
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
