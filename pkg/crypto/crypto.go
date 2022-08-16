package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
	"zmed_patient_manager/pkg/app_errors"
)

type gcmCrypto struct {
	keyCrypto []byte
	keyGcmIV  []byte
	keyAead   []byte
}

func NewGCM(keyCrypto string, keyGcmIV string, keyAead string) (*gcmCrypto, app_errors.AppError) {
	crypto := new(gcmCrypto)
	kc, err := hex.DecodeString(keyCrypto)
	if err != nil {
		return crypto, app_errors.NewInternalServerError("keyCrypto", err)
	}
	kg, err := hex.DecodeString(keyGcmIV)
	if err != nil {
		return crypto, app_errors.NewInternalServerError("keyGcmIV", err)
	}
	ka, err := hex.DecodeString(keyAead)
	if err != nil {
		return crypto, app_errors.NewInternalServerError("keyAead", err)
	}
	crypto.keyGcmIV = kc
	crypto.keyGcmIV = kg
	crypto.keyGcmIV = ka
	return crypto, nil
}

func (c *gcmCrypto) Encrypt(text string) (string, app_errors.AppError) {
	_cipher, err := aes.NewCipher(c.keyCrypto)
	if err != nil {
		return "", app_errors.NewInternalServerError("encrypt", err)
	}
	gcm, err := cipher.NewGCM(_cipher)
	if err != nil {
		return "", app_errors.NewInternalServerError("encrypt", err)
	}
	ciphertext := gcm.Seal(
		nil,
		c.keyGcmIV,
		[]byte(text),
		c.keyAead,
	)
	return c.getEncryptedCiphertextWithTag(ciphertext, gcm), nil
}

func (c *gcmCrypto) Decrypt(ciphertext string) (string, app_errors.AppError) {
	ciphertextEncoded, tagLen, err := c.getCiphertextWithTag(ciphertext)
	if err != nil {
		return "", app_errors.NewInternalServerError("cipher decrypt", err)
	}
	_cipher, err := aes.NewCipher(c.keyCrypto)
	if err != nil {
		return "", app_errors.NewInternalServerError("cipher decrypt", err)
	}
	gcm, err := cipher.NewGCMWithTagSize(_cipher, tagLen)
	if err != nil {
		return "", app_errors.NewInternalServerError("cipher decrypt", err)
	}
	plaintext, err := gcm.Open(
		nil,
		c.keyGcmIV,
		ciphertextEncoded,
		c.keyAead,
	)
	if err != nil {
		return "", app_errors.NewInternalServerError("cipher decrypt", err)
	}
	return string(plaintext), nil
}

func (c *gcmCrypto) getEncryptedCiphertextWithTag(ciphertext []byte, gcm cipher.AEAD) string {
	tagPosition := len(ciphertext) - gcm.Overhead()

	tag := ciphertext[tagPosition:]
	ciphertext = ciphertext[:tagPosition]

	tagEncoded := base64.StdEncoding.EncodeToString(tag)
	ciphertextEncoded := base64.StdEncoding.EncodeToString(ciphertext)

	return tagEncoded + "$" + ciphertextEncoded
}

func (c *gcmCrypto) getCiphertextWithTag(ciphertext string) ([]byte, int, error) {
	if !strings.Contains(ciphertext, "$") {
		return nil, -1, errors.New("invalid ciphertext")
	}

	split := strings.Split(ciphertext, "$")

	tagEncoded, _ := base64.StdEncoding.DecodeString(split[0])
	ciphertextEncoded, _ := base64.StdEncoding.DecodeString(split[1])

	return append(ciphertextEncoded, tagEncoded...), len(tagEncoded), nil
}

func (c *gcmCrypto) stringInSlice(value string, list []string) bool {
	for _, b := range list {
		if b == value {
			return true
		}
	}
	return false
}
