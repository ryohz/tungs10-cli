package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var paddedData []byte = pkcs7Pad(data)
	var encrypted []byte = make([]byte, len(paddedData))
	var cbcEncrypter cipher.BlockMode = cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, paddedData)
	return encrypted, nil
}

func Decrypt(encrypted []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var decrypted []byte = make([]byte, len(encrypted))
	var cbcDecrypter cipher.BlockMode = cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, encrypted)
	return pkcs7unpad(decrypted), nil
}

func pkcs7unpad(data []byte) []byte {
	var dataLength int = len(data)
	var paddingLength int = int(data[dataLength-1])
	return data[:dataLength-paddingLength]
}

func pkcs7Pad(data []byte) []byte {
	var length int = aes.BlockSize - (len(data) % aes.BlockSize)
	var pad []byte = bytes.Repeat([]byte{byte(length)}, length)
	return append(data, pad...)
}

func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}
