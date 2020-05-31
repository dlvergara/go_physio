package tests

import (
	"physiobot/modules/common/db"
	"physiobot/modules/common/logger"
	"io/ioutil"
	"strings"
)

// SetUp data for testing
func SetUp() {
	logger.Info("Setup environment...")
	runSQL("./sql-setup/")
}

// TearDown for testing
func TearDown() {
	logger.Info("Teardown environment...")
	runSQL("./sql-teardown/")
}

func runSQL(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		logger.Fatal(err)
	}

	for _, file := range files {
		logger.Info(file.Name())
		content, err := ioutil.ReadFile(path + file.Name())
		if err != nil {
			logger.Fatal(err)
		}

		queries := strings.Split(string(content), ";")

		for _, query := range queries {
			if strings.TrimSpace(query) != "" {
				if err := db.DbBongo().Exec(query).Error; err != nil {
					logger.Fatal(err)
				}
			}
		}
	}
}
