package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Defina os valores de e e c como big.Int
	e := big.NewInt(3) // Exemplo de expoente
	c := new(big.Int)
	c.SetString("471953743756667994885959002634453703198890516833", 10) // Exemplo de resultado

	// Inicialize i como 1
	m := big.NewInt(1)
	//result := new(big.Int)
	temp := new(big.Int)

	// Dobre i até que i^e seja maior ou igual a c
	for {
		temp.Exp(m, e, nil)
		if temp.Cmp(c) >= 0 {
			fmt.Printf("temp break i: %s\n", temp.String())
			break
		}
		m.Mul(m, big.NewInt(2))
	}

	// Definir o intervalo para busca binária
	low := new(big.Int).Div(m, big.NewInt(2))
	high := new(big.Int).Set(m)

	// Busca binária no intervalo [low, high]
	for low.Cmp(high) <= 0 {
		mid := new(big.Int).Add(low, high)
		mid.Div(mid, big.NewInt(2))

		temp.Exp(mid, e, nil)
		comp := temp.Cmp(c)
		if comp == 0 {
			fmt.Printf("Encontrado i: %s\n", mid.String())
			for i := 0; i < len(mid.String()); i++ {
				
			}
			return
		} else if comp < 0 {
			low.Add(mid, big.NewInt(1))
		} else {
			high.Sub(mid, big.NewInt(1))
		}
	}

	fmt.Println("Não encontrado")
}
