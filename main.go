package main

import (
	"zmed_patient_manager/infrastructure/config"
	"zmed_patient_manager/infrastructure/database"
)

func main() {
	config.PopulateEnv()
	dbConnection := database.InitDBConnection()
	defer database.CloseDBConnection(dbConnection)
}
