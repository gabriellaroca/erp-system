package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"erp-system.com/v1/domain/Security/Encryption/config"
)

func encrypt(text string) (string, error) {
	keyBytes := []byte(config.PrivateKey)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	encrypted := gcm.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}
