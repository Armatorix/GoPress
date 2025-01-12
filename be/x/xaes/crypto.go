package xaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type Crypto struct {
	secret []byte
}

func New(secret []byte) *Crypto {
	return &Crypto{
		secret: secret,
	}
}

// Encrypt encrypts the plain text with the given key. It encodes cipher text with base64.
func (c *Crypto) Encrypt(plaintext string) (string, error) {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("xaes: failed to read secure random bytes: %w", err)
	}

	cipherBlock, err := aes.NewCipher(c.secret)
	if err != nil {
		return "", fmt.Errorf("xaes: failed to create aes cipher: %w", err)
	}

	cfg := cipher.NewCFBEncrypter(cipherBlock, iv)
	cfg.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// MustEncrypt is like Encrypt but panics on error.
func (c *Crypto) MustEncrypt(plaintext string) string {
	ciphertext, err := c.Encrypt(plaintext)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// Decrypt decrypts the cipher text base64 encoded with the given key.
func (c *Crypto) Decrypt(ciphertextb64 string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(string(ciphertextb64))
	if err != nil {
		return "", fmt.Errorf("xaes: failed to decode base64 cipher text: %w", err)
	}

	block, err := aes.NewCipher(c.secret)
	if err != nil {
		return "", fmt.Errorf("xaes: failed to create aes cipher: %w", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("xaes: cipher text too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

// MustDecrypt is like Decrypt but panics on error.
func (c *Crypto) MustDecrypt(key []byte, ciphertextb64 string) string {
	plaintext, err := c.Decrypt(ciphertextb64)
	if err != nil {
		panic(err)
	}
	return plaintext
}
