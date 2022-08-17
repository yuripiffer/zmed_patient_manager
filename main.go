package main

import (
	"fmt"
	"zmed_patient_manager/domain/patient"
	"zmed_patient_manager/infrastructure/config"
	"zmed_patient_manager/infrastructure/database"
	pkgCrypto "zmed_patient_manager/pkg/crypto"
)

func main() {
	config.PopulateEnv()
	dbConnection := database.InitDBConnection()
	defer database.CloseDBConnection(dbConnection)
	crypto, appErr := pkgCrypto.NewGCM(config.ENV.KeyCrypto, config.ENV.KeyAead, config.ENV.KeyGcmIV)
	if appErr != nil {
		panic(appErr)
	}
	patientUseCase := patient.New(dbConnection, crypto)

	fmt.Println(patientUseCase)
}
