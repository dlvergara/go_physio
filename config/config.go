package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Configuration struct {
	ExtDomains struct {
		Ftn    string
	} `yaml:"extDomains"`
	Redis      struct {
		Host           string
		Port           string
		Password       string
		Db             int
		ExpirationTime int `yaml:"expirationTime"`
	}
	Db struct {
		LukexDb struct {
			Adapter   string
			Database  string
			Username  string
			Password  string
			Host      string
			Port      string
			IdleConns int `yaml:"idleConns"`
			OpenConns int `yaml:"openConns"`
		} `yaml:"lukexDb"`
	} `yaml:"db"`
	Jwt struct {
		Alg string
		Key string
		Exp int64
	}
	App struct {
		Port string
	}
	LogEnabled bool `yaml:"logEnabled"`
}

var configuration Configuration

func init() {
	appPath, _ := os.Getwd()
	println(appPath);
	configAppPath := appPath + "/config/global.yaml"

	if fileExists(appPath + "/config/local.yaml") {
		fmt.Println("Local config exists")
		configAppPath = appPath + "/config/local.yaml"
	}

	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&configuration, configAppPath)
	if err != nil {
		fmt.Println(err)
	}
	//configuration.Db.LukexDb.Host = configuration.Db.LukexDb.Host
	//configuration.Db.LukexDb.Port = configuration.Db.LukexDb.Port
	//configuration.Db.LukexDb.Database = configuration.Db.LukexDb.Database
	//configuration.Db.LukexDb.Username = configuration.Db.LukexDb.Username
	//configuration.Db.LukexDb.Password = configuration.Db.LukexDb.Password
	/*
	configuration.Db.Bongo.Host = getenv("DB_BONGO_HOST", configuration.Db.Bongo.Host)
	configuration.Db.Bongo.Port = getenv("DB_BONGO_PORT", configuration.Db.Bongo.Port)
	configuration.Db.Bongo.Database = getenv("DB_BONGO_DATABASE", configuration.Db.Bongo.Database)
	configuration.Db.Bongo.Username = getenv("DB_BONGO_USERNAME", configuration.Db.Bongo.Username)
	configuration.Db.Bongo.Password = getenv("DB_BONGO_PASSWORD", configuration.Db.Bongo.Password)
	configuration.Db.Masterbook.Host = getenv("DB_BONGO_HOST", configuration.Db.Masterbook.Host)
	configuration.Db.Masterbook.Port = getenv("DB_MASTERBOOK_PORT", configuration.Db.Masterbook.Port)
	configuration.Db.Masterbook.Database = getenv("DB_MASTERBOOK_DATABASE", configuration.Db.Masterbook.Database)
	configuration.Db.Masterbook.Username = getenv("DB_MASTERBOOK_USERNAME", configuration.Db.Masterbook.Username)
	configuration.Db.Masterbook.Password = getenv("DB_MASTERBOOK_PASSWORD", configuration.Db.Masterbook.Password)
	configuration.Redis.Host = getenv("REDIS_HOST", configuration.Redis.Host)
	configuration.Redis.Port = getenv("REDIS_PORT", configuration.Redis.Port)
	configuration.Redis.Password = getenv("REDIS_PASSWORD", configuration.Redis.Password)
	configuration.App.Port = getenv("APP_PORT", configuration.App.Port)
	configuration.FcbDomains.Secure = getenv("SECURE_DOMAIN", configuration.FcbDomains.Secure)
	configuration.FcbDomains.Tools = getenv("TOOLS_DOMAIN", configuration.FcbDomains.Tools)
	configuration.FcbDomains.Services = getenv("SERVICES_DOMAIN", configuration.FcbDomains.Services)
	configuration.ExtDomains.Ftn = getenv("FTN_DOMAIN", configuration.ExtDomains.Ftn)
	configuration.Jwt.Key = getenv("JWT_KEY", configuration.Jwt.Key)
	*/
}

// GetConfig return global.yaml config on a Configuration struct
func GetConfig() Configuration {
	return configuration
}

func getenv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
