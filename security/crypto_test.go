package security

import (
	"encoding/hex"
	"testing"
)

func TestSha256(t *testing.T) {
	str := []byte("i am apple")
	result := Sha256(str)
	ans := "310c6e3de38bdc9bc74db2db36d0d0c26bbe4185804c4f74aead51991fb1bb70"
	if hex.EncodeToString(result) != ans {
		t.Error("hash result incorrect")
	}
}

func TestCrypto(t *testing.T) {
	plain := "i am apple, you are orange"
	key := []byte(RandStr(32))
	iv := []byte(RandStr(16))
	crypted := AESCBCPK5Encrypt([]byte(plain), key, iv)
	decrypted := AESCBCPK5Decrypt(crypted, key, iv)
	if string(decrypted) != plain {
		t.Error("encrypt/decrypt result incorrect")
	}
}
