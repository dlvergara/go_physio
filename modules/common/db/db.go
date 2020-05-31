package db

import (
	"physiobot/config"
	"physiobot/modules/common/logger"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Handler struct to keep db conections
type Handler struct {
	Bongo      *gorm.DB
	Masterbook *gorm.DB
}

type dbConf struct {
	Adapter   string
	Database  string
	Username  string
	Password  string
	Host      string
	Port      string
	IdleConns int
	OpenConns int
}

var dbHandler *Handler

// Init conections with config from global.yaml
func init() {
	//configuration := config.GetConfig()
	/*
	bongoConfig := dbConf{
		Adapter:   configuration.Db.Bongo.Adapter,
		Database:  configuration.Db.Bongo.Database,
		Username:  configuration.Db.Bongo.Username,
		Password:  configuration.Db.Bongo.Password,
		Host:      configuration.Db.Bongo.Host,
		Port:      configuration.Db.Bongo.Port,
		IdleConns: configuration.Db.Bongo.IdleConns,
		OpenConns: configuration.Db.Bongo.OpenConns,
	}

	masterbookConfig := dbConf{
		Adapter:   configuration.Db.Masterbook.Adapter,
		Database:  configuration.Db.Masterbook.Database,
		Username:  configuration.Db.Masterbook.Username,
		Password:  configuration.Db.Masterbook.Password,
		Host:      configuration.Db.Masterbook.Host,
		Port:      configuration.Db.Masterbook.Port,
		IdleConns: configuration.Db.Masterbook.IdleConns,
		OpenConns: configuration.Db.Masterbook.OpenConns,
	}

	dbHandler = new(Handler)
	dbHandler.Bongo = mysqlConn(bongoConfig)
	dbHandler.Masterbook = mysqlConn(masterbookConfig)
	*/
}

// Manager return dbHandler instance
func Manager() *Handler {
	return dbHandler
}

func DbBongo() *gorm.DB {
	return dbHandler.Bongo
}

func mysqlConn(dbEnv dbConf) *gorm.DB {
	var (
		db               *gorm.DB
		connectionString string
		err              error
	)

	username := dbEnv.Username
	password := dbEnv.Password
	host := dbEnv.Host
	port := dbEnv.Port
	configuration := config.GetConfig()

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password,
		host, port,
		dbEnv.Database)

	logger.Info(connectionString)

	if db, err = gorm.Open("mysql", connectionString); err != nil {
		logger.Error("Error open")
		logger.Fatal(err.Error())
		return db
	}
	if err = db.DB().Ping(); err != nil {
		logger.Error("Error ping")
		logger.Fatal(err.Error())
		return db
	}

	db.LogMode(configuration.LogEnabled)
	db.DB().SetMaxIdleConns(dbEnv.IdleConns)
	db.DB().SetMaxOpenConns(dbEnv.OpenConns)

	logger.Success("Conected to " + dbEnv.Database)
	return db
}
