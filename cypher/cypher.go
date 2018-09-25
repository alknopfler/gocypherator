package cypher


import (
	"crypto/aes"
	"crypto/cipher"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	math_rand "math/rand"

	"github.com/alknopfler/Gologger/gologger"
)

type KeyCypher struct {
	Key 	[]byte
}

func InitKeyCypher(cypherType int) *KeyCypher {
	var k = KeyCypher{}
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,:;.<>"
	if cypherType != 16 && cypherType != 24 && cypherType != 32{
		return &k
	}
	b := make([]byte, cypherType)
	for i := range b {
		b[i] = letterBytes[math_rand.Intn(len(letterBytes))]
	}
	k.Key = b
	return &k
}

func (k *KeyCypher) EncryptString(text string) string {
	plaintext := []byte(text)
	block, err := aes.NewCipher(k.Key)
	if err != nil {
		gologger.Print("ERROR", 8, "Error while cyphering key", "cypher.go")
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(crypto_rand.Reader, iv); err != nil {
		gologger.Print("ERROR", 8, "Error while encrypting key", "cypher.go")
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func (k *KeyCypher) DecryptString(cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher(k.Key)
	if err != nil {
		gologger.Print("ERROR", 8, "Error while cyphering", "cipher.go")
	}
	if len(ciphertext) < aes.BlockSize {
		gologger.Print("ERROR", 8, "Ciphertext too short", "cipher.go")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}


