package main

import (
	"fmt"
	"time"
	"zmed_patient_manager/infrastructure/config"
	"zmed_patient_manager/infrastructure/database"
)

func main() {
	fmt.Println(time.Now().Year())
	config.PopulateEnv()
	dbConnection := database.InitDBConnection()
	defer database.CloseDBConnection(dbConnection)
}
