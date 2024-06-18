package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Função para decodificar a string de números
func decodeString(numericalString string) string {
	var decodedString strings.Builder
	i := 0
	for i < len(numericalString) {
		// Tentativa de pegar dois dígitos
		twoDigits, err := strconv.Atoi(numericalString[i : i+2])
		if err != nil {
			fmt.Println("Erro de conversão:", err)
			return ""
		}
		// Se os dois dígitos forem um caractere válido
		if twoDigits >= 65 && twoDigits <= 90 {
			decodedString.WriteByte(byte(twoDigits))
			i += 2
			continue
		}

		// Tentativa de pegar três dígitos
		threeDigits, err := strconv.Atoi(numericalString[i : i+3])
		if err != nil {
			fmt.Println("Erro de conversão:", err)
			return ""
		}
		// Se os três dígitos forem um caractere válido
		if threeDigits >= 65 && threeDigits <= 122 {
			decodedString.WriteByte(byte(threeDigits))
			i += 3
		} else {
			fmt.Println("Erro: código ASCII inválido encontrado.")
			return ""
		}
	}
	return decodedString.String()
}

func main() {
	string_numer := "7785738479666977"
	result := decodeString(string_numer)
	fmt.Println("Mensagem decodificada:", result)
}
