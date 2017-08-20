package logic

import "gitlab.pec.ir/cloud/sync-service/models"
import "gitlab.pec.ir/cloud/sync-service/db"

const fILEADDRESS = "/config.json"

//LoadConfig loads config from json file
//retunrs subscriber model
func LoadConfig() []models.Subscriber {
	return db.LoadConfiguration(fILEADDRESS)
}
