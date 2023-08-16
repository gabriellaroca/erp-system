package config

import "os"

var SecretJwt = os.Getenv("SECRET_JWT")
