package calculadora

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

func Restar(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}

func DividirSeguro(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, fmt.Errorf("no se puede dividir por cero")
}

func Dividir(a, b int) int {
	return a / b
}
