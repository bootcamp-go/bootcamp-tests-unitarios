package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumarConTestify(t *testing.T) {
	//arrange
	a, b := 3, 5
	resultadoEsperado := 8

	//act
	resultado := Sumar(a, b)

	//assert
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestRestarConTestify(t *testing.T) {
	//arrange
	a, b := 5, 3
	resultadoEsperado := 2

	//act
	resultado := Restar(a, b)

	//assert
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestDividirSeguroConTestify(t *testing.T) {

	t.Run("6 dividido 3 debe dar 2, y error nulo", func(t *testing.T) {
		//arrange
		a, b := 6, 3
		resultadoEsperado := 2

		//act
		resultado, err := DividirSeguro(a, b)

		//assert
		assert.Nil(t, err)
		assert.Equal(t, resultadoEsperado, resultado)

	})

	t.Run("6 dividido 0 debe dar error no nulo y zero value de int", func(t *testing.T) {
		//arrange
		a, b := 6, 0
		resultadoEsperado := 0

		//act
		resultado, err := DividirSeguro(a, b)

		//assert
		assert.NotNil(t, err)
		assert.Equal(t, resultadoEsperado, resultado)

	})
}

func TestDividirConTestify(t *testing.T) {
	// Que pasa si ocurre un panic durante la ejecución del test?

	//arrange
	a, b := 5, 0
	resultadoEsperado := 0

	//act
	resultado := Dividir(a, b)

	//assert
	assert.Equal(t, resultadoEsperado, resultado)
}
