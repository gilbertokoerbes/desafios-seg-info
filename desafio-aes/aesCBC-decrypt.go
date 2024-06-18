package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func addPKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	fmt.Println("Padding => ", padding, padText)
	return append(data, padText...)
}
func main() {

	/////////////////////////////////////////////////////////////////////
	/////////////////////// D E C R Y P T ///////////////////////////////
	/////////////////////////////////////////////////////////////////////
	key, _ := hex.DecodeString("140b41b22a29beb4061bda66b6747e14")

	//MENSAGEM 1 RECEBIDA
	//ciphertext, _ := hex.DecodeString("9752C89280BC97E0E66CE51688E76C5F7FB10ADF452EE59DDBF336499EFDCB84F7D3117E24AA1F0EC2D044EB0E520CCACA0A4747E4D04F271F3F51C7A34FEFFE57931116C9715823DDB927ECA6913F530C08FDE0AFC361098A80BA21FD2F3F151E12ACF82DD2EE87B000C132A35FC7EC")

	//MENSAGEM 2 RECEBIDA
	ciphertext, _ := hex.DecodeString("d2790dea5adb64f3f6efcc43d2ab428adf0f5fa12cce65203a5156ffa0564eca95772032006e058c738adc8d3196e72a89f0d5494f106d3eaced5863f2efbaab64666e00b0ba21e62bca5a278abcaeebca76929b12e4c96f21415c4d890e95df7da644769037484a9d4e2d3d1dda6ec28ca406f9234aa1f86e33d03968d7b439f8384f0c706ae4d0f3d0fad2510747f5")

	//Block, em Golang, define um objeto BlockMode com base na chave, que define os atributos como quantidade de rounds e tamanho de chave
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// if len(ciphertext) < aes.BlockSize {
	// 	panic("ciphertext too short")
	// }

	iv := ciphertext[:aes.BlockSize]        //pegar os primeiros bytes que representam IV
	ciphertext = ciphertext[aes.BlockSize:] //pegar os demais bytes que representam mensagem
	//ciphertext = addPKCS7Padding(ciphertext, aes.BlockSize)

	fmt.Println("aes.BlockSize", aes.BlockSize)
	fmt.Println("iv", iv)
	fmt.Println("iv hex", hex.EncodeToString(iv))
	fmt.Println("Mensagem cifrada recebida", ciphertext)
	fmt.Println("Length Mensagem recebida", len(ciphertext))

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	//Aqui criamos o modo de operação, passando o objeto Block(informações da chave, tamanho de chave e rounds) e o IV
	mode := cipher.NewCBCDecrypter(block, iv)

	deciphertext := ciphertext
	mode.CryptBlocks(deciphertext, ciphertext) //com o modo definido - parametro1: valor onde a mensagem decifrada é gravada, parametro2: mensagem cifrada

	fmt.Println("%s\n Mensagem decifrada recebida em hexa=> ", (deciphertext[:]))
	fmt.Println("%s\n Mensagem decifrada recebida => ", string(deciphertext[:]))
	fmt.Println("Length Mensagem decifrada recebida => ", len(deciphertext))

}
