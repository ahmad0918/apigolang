package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/viper"
)

func EncryptAES(value string) string {
	c, err := aes.NewCipher([]byte(GetSalt()))
	if err != nil {
		return err.Error()
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return fmt.Sprintf("%x", gcm.Seal(nonce, nonce, []byte(value), nil))
}

func DecryptAES(value string) string {
	cipherText, err := hex.DecodeString(value)
	if err != nil {
		fmt.Printf("error decrypt: %v", err)
		return ""
	}
	c, err := aes.NewCipher([]byte(GetSalt()))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return ""
	}

	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("%s", plaintext)

}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GetSalt() string {

	type TurningCheck struct {
		Check string `json:"check"`
	}

	var check TurningCheck

	encSalt := viper.GetString("turning_check")
	decoded, _ := base64.StdEncoding.DecodeString(encSalt)
	decodedCheck := reverse(string(decoded))
	resultString, _ := base64.StdEncoding.DecodeString(decodedCheck)

	err := json.Unmarshal(resultString, &check)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := check.Check

	return result
}
