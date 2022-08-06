package calculadora

import (
	"testing"
)

func TestSumar(t *testing.T) {
	//arrange
	a, b := 3, 5
	resultadoEsperado := 8

	//act
	resultado := Sumar(a, b)

	//assert
	if resultadoEsperado != resultado {
		t.Errorf("la funcion sumar, con los argumentos a=%d, b=%d devolvió %d pero se esperaba %d", a, b, resultado, resultadoEsperado)
	}
}

func TestRestar(t *testing.T) {
	//arrange
	a, b := 5, 3
	resultadoEsperado := 8

	//act
	resultado := Restar(a, b)

	//assert
	if resultadoEsperado != resultado {
		t.Errorf("la funcion restar, con los argumentos a=%d, b=%d devolvió %d pero se esperaba %d", a, b, resultado, resultadoEsperado)
	}
}

func TestDividirSeguro(t *testing.T) {

	t.Run("6 dividido 3 debe dar 2, y error nulo", func(t *testing.T) {
		//arrange
		a, b := 6, 3
		resultadoEsperado := 2

		//act
		resultado, err := DividirSeguro(a, b)

		//assert
		if err != nil {
			t.Errorf("la funcion DividirSeguro, con los argumentos a=%d, b=%d devolvió un error no nil pero se esperaba que sea nil", a, b)
		}
		if resultadoEsperado != resultado {
			t.Errorf("la funcion DividirSeguro, con los argumentos a=%d, b=%d devolvió %d pero se esperaba %d", a, b, resultado, resultadoEsperado)
		}

	})

	t.Run("6 dividido 0 debe dar error no nulo y zero value de int", func(t *testing.T) {
		//arrange
		a, b := 6, 0
		resultadoEsperado := 0

		//act
		resultado, err := DividirSeguro(a, b)

		//assert
		if err == nil {
			t.Errorf("la funcion DividirSeguro, con los argumentos a=%d, b=%d devolvió un error nil pero se esperaba que sea no nil", a, b)
		}
		if resultadoEsperado != resultado {
			t.Errorf("la funcion DividirSeguro, con los argumentos a=%d, b=%d devolvió %d pero se esperaba %d", a, b, resultado, resultadoEsperado)
		}

	})
}
