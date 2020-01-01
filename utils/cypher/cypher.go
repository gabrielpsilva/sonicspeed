package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)


// Encrypt encrypts something using package crypto/aes
func Encrypt(plaintext []byte, keyString string) []byte {

	key := []byte(keyString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return cipherText
}

// Decrypt decrypts something using package crypto/aes
func Decrypt(encrypted []byte, keyString string) []byte {

	key := []byte(keyString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(encrypted) < aes.BlockSize {
		panic("Text is too short")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return encrypted
}

func DecryptFile(filePath, keyString string) ([]byte, error) {


	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return Decrypt(dat, keyString), nil
}

/*
// EncryptAndBase64 will cipher a readable and return in base64
func EncryptAndBase64(plainstring, keystring string) string {

	data := Encrypt(plainstring, keystring)
	return b64.StdEncoding.EncodeToString(data)
}

// Base64AndDecrypt will make a cipher in base64 readable
func Base64AndDecrypt(b64string, keystring string) string {

	data, err := b64.StdEncoding.DecodeString(b64string)
	if err != nil {
		panic(err)
	}
	return Decrypt(data, keystring)

}



*/

