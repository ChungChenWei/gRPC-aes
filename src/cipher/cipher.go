package cipher

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/zenazn/pkcs7pad"
)

func Encrypt(plaintext []byte, key string, iv string) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	padded := pkcs7pad.Pad(plaintext, block.BlockSize())
	encrypted := make([]byte, len(padded))

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	blockMode.CryptBlocks(encrypted, padded)

	return encrypted
}

func Decrypt(encrypted []byte, key string, iv string) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	decrypted := make([]byte, len(encrypted))

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	blockMode.CryptBlocks(decrypted, encrypted)

	plaintext, err := pkcs7pad.Unpad(decrypted)
	if err != nil {
		panic(err)
	}

	return plaintext
}
