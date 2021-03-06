package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func encryptMessage(message string) string {
	plaintext := []byte(message)

	// aes encryption string
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"

	fmt.Println(len(key_text))

	// Create the aes encryption algorithm
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	// Encrypted string
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	// fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// Decrypt strings
	// cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	// plaintextCopy := make([]byte, len(plaintext))
	// cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	// fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

	return fmt.Sprintf("%x", ciphertext)
}
