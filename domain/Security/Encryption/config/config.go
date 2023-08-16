package config

import "os"

var PrivateKey string

func SetPrivateKey() (PrivateKey string) {
	PrivateKey = os.Getenv("PRIVATE_KEY")
	return
}
