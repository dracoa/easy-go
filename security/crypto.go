package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"math/rand"
	"strings"
	"time"
)

func Sha256(sec []byte) []byte {
	h := sha256.New()
	h.Write(sec)
	return h.Sum(nil)
}

func AESCBCPK5Encrypt(src []byte, key []byte, iv []byte) []byte {
	if len(src) < 1 {
		panic("plain content empty")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)
	content := PKCS5Padding(src, block.BlockSize())
	crypted := make([]byte, len(content))
	encrypter.CryptBlocks(crypted, content)
	return crypted
}

func AESCBCPK5Decrypt(src []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	decrypter := cipher.NewCBCDecrypter(block, iv)
	content := make([]byte, len(src))
	decrypter.CryptBlocks(content, src)
	return PKCS5Trimming(content)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func RandStr(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
