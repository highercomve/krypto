package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"strconv"
)

// Crypter create Crypter
type Crypter struct {
	key   []byte
	iv    []byte
	block cipher.Block
}

// NewCrypter define new crypter
func NewCrypter(key []byte, iv []byte) (*Crypter, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, errors.New("IV length has to be equal to: " + strconv.Itoa(blockSize))
	}

	return &Crypter{key, iv, block}, nil
}

// Encrypt Encrypt
func (c *Crypter) Encrypt(plainText []byte) ([]byte, error) {
	origData := pkcs5Padding(plainText, c.block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(c.block, c.iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

// Decrypt decrypt
func (c *Crypter) Decrypt(cipherText []byte) ([]byte, error) {
	blockMode := cipher.NewCBCDecrypter(c.block, c.iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
