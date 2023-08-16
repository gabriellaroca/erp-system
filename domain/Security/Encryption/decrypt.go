package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"erp-system.com/v1/domain/Security/Encryption/config"
)

func decrypt(encryptedText string) (string, error) {
	keyBytes := []byte(config.PrivateKey)
	encrypted, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return "", fmt.Errorf("texto criptografado inválido")
	}

	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
