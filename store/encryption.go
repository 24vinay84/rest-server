// Golang AES Encryption/Decryption

package store

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "io"
)


var key = []byte("a very very very very secret key")
var null = ""

func encrypt(str string) string {

	text := []byte(str)
    block, err := aes.NewCipher(key)
    if err != nil {
        return null
    }
    b := base64.StdEncoding.EncodeToString(text)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return null
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	myString := string(ciphertext[:])
    return myString
}

func decrypt(str string) string {
    
	text := []byte(str)
	block, err := aes.NewCipher(key)
    if err != nil {
        return null
    }
    if len(text) < aes.BlockSize {
        return null
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return null
    }
	myString := string(data)
    return myString
}