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
	Lukex      *gorm.DB
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
	configuration := config.GetConfig()

	lukexConfig := dbConf{
		Adapter:   configuration.Db.LukexDb.Adapter,
		Database:  configuration.Db.LukexDb.Database,
		Username:  configuration.Db.LukexDb.Username,
		Password:  configuration.Db.LukexDb.Password,
		Host:      configuration.Db.LukexDb.Host,
		Port:      configuration.Db.LukexDb.Port,
		IdleConns: configuration.Db.LukexDb.IdleConns,
		OpenConns: configuration.Db.LukexDb.OpenConns,
	}

	dbHandler = new(Handler)
	dbHandler.Lukex = mysqlConn(lukexConfig)
}

// Manager return dbHandler instance
func Manager() *Handler {
	return dbHandler
}

func DbLukex() *gorm.DB {
	return dbHandler.Lukex
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
