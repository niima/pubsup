package db

import (
	"encoding/json"
	"fmt"
	"os"

	"gitlab.pec.ir/cloud/sync-service/models"
)

//LoadConfiguration loads config file
func LoadConfiguration(file string) []models.Subscriber {
	var config []models.Subscriber
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
