package api

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
)

func CompressString(input string) string {
	key := []byte("90319283-129-3091-39-129") // Задайте свой ключ шифрования здесь
	plaintext := []byte(input)

	// Инициализация блочного шифра AES
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// Генерация IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	// Создание галочки для шифрования
	stream := cipher.NewCTR(block, iv)

	// Шифрование строки
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// Возвращение сжатой строки в формате base64
	return base64.StdEncoding.EncodeToString(ciphertext)[:6]
}
