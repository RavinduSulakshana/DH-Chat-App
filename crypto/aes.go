package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// create cyphertext slice with space for IV+encrypted data
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	//Generate random iv and fill it
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// decrypt ciphertext using aes counter mode
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	if len(ciphertext) < len(aes.BlockSize) {
		return nil, errors.New("ciphertext is very short")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//extract IV and encrypted message from ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext := ciphertext[aes.BlockSize:]

	//decrypt using CTR mode
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}
