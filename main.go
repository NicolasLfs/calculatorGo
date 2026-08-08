package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 float64
	var numero2 float64
	var operador string
	var continuar string

	for {
		fmt.Println("Digite o primeiro valor:")
		fmt.Scan(&numero1)
		fmt.Println("Digite o operador:")
		fmt.Scan(&operador)
		fmt.Println("Digite o segundo valor:")
		fmt.Scan(&numero2)

		resultado, err := calcular(numero1, numero2, operador)

		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		fmt.Println("Resultado:", resultado)

		fmt.Println("Deseja fazer outro calculo? (s/n)")
		fmt.Scan(&continuar)

		if continuar != "s" {
			fmt.Println("Até logo!")
			break
		}
	}
}

func calcular(n1 float64, n2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, errors.New("Não é possivel dividir por zero")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("Operador Invalido")
	}
}
