package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func addPKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	fmt.Println("Padding => ", padding, padText)
	return append(data, padText...)
}
func main() {

	/////////////////////////////////////////////////////////////////////
	/////////////////////// E N C R Y P T ///////////////////////////////
	/////////////////////////////////////////////////////////////////////

	key, _ := hex.DecodeString("140b41b22a29beb4061bda66b6747e14")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	message := "4e657874205468757273646179206f6e65206f66207468652062657374207465616d7320696e2074686520776f726c642077696c6c2066616365206120626967206368616c6c656e676520696e20746865204c696265727461646f72657320646120416d6572696361204368616d70696f6e736869702e"

	// Inverter a mensagem em texto plano
	// r_message := []rune(message)
	// for i, j := 0, len(r_message)-1; i < j; i, j = i+1, j-1 {
	// 	r_message[i], r_message[j] = r_message[j], r_message[i]
	// }

	message_bytes, _ := hex.DecodeString(message)

	fmt.Println("Mensagem:", (string(message_bytes)))
	fmt.Println("\nlen message_bytes", len(string(message_bytes)))

	message_bytes_with_padding := addPKCS7Padding(message_bytes, aes.BlockSize)
	fmt.Println("\n message_bytes_with_padding", message_bytes_with_padding)
	if len(string(message_bytes_with_padding))%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	//generate new IV random
	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("iv = make([]byte, 16)", iv)

	ciphertext_to_send := make([]byte, aes.BlockSize+len(message_bytes_with_padding))
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext_to_send[aes.BlockSize:], message_bytes_with_padding)

	fmt.Println("ciphertext_to_send message_bytes sem iv = ", hex.EncodeToString(ciphertext_to_send))

	//Add IV in ciphertext_to_send
	for i := 0; i < aes.BlockSize; i++ {
		ciphertext_to_send[i] = iv[i]

	}

	fmt.Println("ciphertext_to_send message_bytes = ", hex.EncodeToString(ciphertext_to_send))

}
