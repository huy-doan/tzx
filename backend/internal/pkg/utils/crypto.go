package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// EncryptAESGCM AES-GCM方式で暗号化を行う
// key長は16byte(AES-128) または32byte(AES-256)である必要がある
// https://pkg.go.dev/crypto/cipher#NewGCM
// https://qiita.com/ken5scal/items/da387d6db8b8308474e6
func EncryptAESGCM(key []byte, plain []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize()) // Unique nonce is required(NonceSize 12byte)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nil, nonce, plain, nil)
	encrypted = append(nonce, encrypted...)

	return encrypted, nil
}

// DecryptAESGCM AES-GCM方式で復号化を行う
// key長は16byte(AES-128) または32byte(AES-256)である必要がある
// https://pkg.go.dev/crypto/cipher#NewGCM
// https://qiita.com/ken5scal/items/da387d6db8b8308474e6
func DecryptAESGCM(key []byte, encrypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte(""), err
	}

	if len(encrypted) < gcm.NonceSize() {
		return []byte(""), fmt.Errorf("error: slice bounds out of range encrypted size: %d key size: %d", len(encrypted), gcm.NonceSize())
	}

	nonce := encrypted[:gcm.NonceSize()]
	plain, err := gcm.Open(nil, nonce, encrypted[gcm.NonceSize():], nil)
	if err != nil {
		return []byte(""), err
	}

	return plain, nil
}

func EncryptAESGCMString(key string, value string) (string, error) {
	if value == "" {
		return "", nil
	}

	if key == "" {
		return "", fmt.Errorf("key is empty")
	}

	keyDecode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	decryptStringByte, err := EncryptAESGCM(keyDecode, []byte(value))
	if err != nil {
		return "", err
	}

	decryptString := base64.URLEncoding.EncodeToString(decryptStringByte)
	return decryptString, nil
}

func DecryptAESGCMString(key string, value string) (string, error) {
	if value == "" {
		return "", nil
	}

	if key == "" {
		return "", fmt.Errorf("key is empty")
	}

	keyDecode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	valueDecode, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	decryptStringByte, err := DecryptAESGCM(keyDecode, valueDecode)
	if err != nil {
		return "", err
	}

	decryptString := string(decryptStringByte)
	return decryptString, nil
}
