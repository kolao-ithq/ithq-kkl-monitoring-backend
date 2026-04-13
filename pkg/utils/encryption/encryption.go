package encryption

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	CipherKey = []byte("thisis32bitlongpassphraseimusing")
)

func Encrypt(code string) ([]byte, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(code), 14)
	if err != nil {
		return nil, err
	}
	return encrypt, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func EncryptAES(plaintext string) (string, error) {

	c, err := aes.NewCipher(CipherKey)
	if err != nil {
		return "", err
	}

	out := make([]byte, 32)
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out), nil
}
func DecryptAES(ct string) (string, error) {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(CipherKey)
	if err != nil {
		return "", err
	}
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
	return s, nil
}

func EncodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return strings.TrimSpace(string(base64Text))
}
