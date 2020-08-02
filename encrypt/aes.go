/**
aes加/解密
加密模式 CBC
*/
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// AesCBCEncrypt aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
// https://studygolang.com/articles/14251?fr=sidebar
func AesCBCEncrypt(rawData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	rawData = pkcs7Padding(rawData, blockSize)

	var cipherText []byte
	if len(iv) > 0 {
		// 固定向量
		cipherText = make([]byte, len(rawData))
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(cipherText, rawData)

	} else {
		// 随机向量
		cipherText = make([]byte, blockSize+len(rawData))
		iv = cipherText[:blockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}

		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(cipherText[blockSize:], rawData)
	}

	return cipherText, nil
}

// AesCBCDecrypt
func AesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	if len(iv) <= 0 {
		// 向量参数为空时，取随机向量
		iv = encryptData[:blockSize]
		encryptData = encryptData[blockSize:]
	}

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	return pkcs7UnPadding(encryptData)
}

func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs7UnPadding(originData []byte) ([]byte, error) {
	length := len(originData)
	unpadding := int(originData[length-1])
	if length < unpadding {
		return nil, errors.New("ciphertext is not a valid size")
	}
	return originData[:(length - unpadding)], nil
}
