package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("e46cedc3752575e8be3d52809a654565")
	//plaintext, _ := hex.DecodeString("5468697320697320612073656e74656e636520746f20626520656e63727970746564207573696e672041455320616e6420435452206d6f64652e")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext, _ := hex.DecodeString("010101010101010101010101010101012C733DE403A769B9CB6D72E9E323B9F0F5FEAE7D748D1B")

	fmt.Printf("\n ciphertext hex: \n", hex.EncodeToString(ciphertext))

	iv := ciphertext[:aes.BlockSize]
	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	panic(err)
	// }

	//stream := cipher.NewCTR(block, iv)
	//stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("\n ciphertext hex: \n", hex.EncodeToString(ciphertext))
	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertext with NewCTR.

	plaintext2 := make([]byte, (len(ciphertext) - aes.BlockSize))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("\n plaintext2 %s", string(plaintext2))
}
