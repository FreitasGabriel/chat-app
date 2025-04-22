package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/gorilla/websocket"
)

func ReadMessageFromWebscoket(ws *websocket.Conn, broadcast chan entity.Message, cypherKey []byte) {
	for {
		var message entity.Message
		if err := ws.ReadJSON(&message); err != nil {
			logger.Error("Error to read message", err)
			return
		}


		message.ID = entity.NewUUID().String()
		message.CreatedAt = time.Now()

		encryptedMsg := EncryptMessageByAES(cypherKey, message.Message)
		message.Message = encryptedMsg
		fmt.Println("message", message)
		broadcast <- message
	}
}

func WriteMessageOnWebsocket(broadcast chan entity.Message, clients map[*websocket.Conn]bool, cypherKey []byte) {
	for {
		msg := <-broadcast
		decryptedMSG := DecryptMessageByAES(cypherKey, msg.Message)
		for client := range clients {
			msg.Message = decryptedMSG
			err := client.WriteJSON(msg)
			if err != nil {
				logger.Error("error to handle message", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func EncryptMessageByAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("error to create cypher to encrypt message", err)
		return ""
	}

	padding := c.BlockSize() - len(plaintext)%c.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	plaintextBytes := []byte(plaintext)
	plaintextBytes = append(plaintextBytes, padtext...)

	//allocate space for ciphered data
	out := make([]byte, len(plaintextBytes))

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		logger.Error("error generating IV", err)
		return "" // Handle the error appropriately
	}

	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(out, []byte(plaintextBytes))

	combined := append(iv, out...)

	return hex.EncodeToString(combined)
}

func DecryptMessageByAES(key []byte, ct string) string {
	ciphertext, err := hex.DecodeString(ct)
	if err != nil {
		logger.Error("error to decode ciphertext", err)
		return ""
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("error to create cypher to decrypt message", err)
		return ""
	}

	if len(ciphertext) < aes.BlockSize { // Check before slicing for IV!
		logger.Error("ciphertext too short", nil)
		return ""
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Check ciphertext length *after* removing IV
	if len(ciphertext) == 0 {
		logger.Error("ciphertext too short after removing IV", nil)
		return ""
	}

	if len(ciphertext)%c.BlockSize() != 0 {
		logger.Error("ciphertext is not a multiple of the block size", nil) // Correct error handling
		return ""
	}

	mode := cipher.NewCBCDecrypter(c, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	if len(ciphertext) == 0 {
		logger.Error("ciphertext is empty after decryption", nil)
		return ""
	}

	padding := int(ciphertext[len(ciphertext)-1])

	// Check valid padding value
	if padding < 1 || padding > aes.BlockSize {
		logger.Error("invalid padding value", nil)
		return "" // Or handle the error as needed
	}

	return string(ciphertext)
}
