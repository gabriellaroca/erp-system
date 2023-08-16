package main

import (
	"os"

	mongoConfig "erp-system.com/v1/adapter/mongodb/config"
	api "erp-system.com/v1/api/config"
	cryptConfig "erp-system.com/v1/domain/Security/Encryption/config"
	"github.com/joho/godotenv"
)

func main() {

	mongoConfig.SetMongoClient()

	godotenv.Load(os.Getenv("ENVIRONMENT") + ".env")

	cryptConfig.SetPrivateKey()

	api.Setup()
}
