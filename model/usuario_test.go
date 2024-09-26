package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidaUsuario(t *testing.T) {
	t.Run("Deve retornar nulo quando o cadastro do usuário estiver completo", func(t *testing.T) {
		// arange
		usuario := Usuario{Nome: "fulano", Email: "ashd", Senha: "senha"}

		// act
		err := ValidaUsuario(usuario)

		// assert
		assert.NoError(t, err)
	})

	t.Run("Deve retornar erro quando usuário não tiver senha", func(t *testing.T) {
		// arange
		usuario := Usuario{Nome: "fulano", Email: "ashd", Senha: ""}
		expectedErrorMessage := "Senha do usuário deve ser preenchida"

		// act
		err := ValidaUsuario(usuario)

		// assert
		assert.EqualError(t, err, expectedErrorMessage)
	})

	t.Run("Deve retornar erro quando usuário não tiver nome", func(t *testing.T) {
		// arange
		usuario := Usuario{Nome: "", Email: "ashd", Senha: "sdasd"}
		expectedErrorMessage := "Nome do usuário deve ser preenchida"

		// act
		err := ValidaUsuario(usuario)

		// assert
		assert.EqualError(t, err, expectedErrorMessage)
	})

	t.Run("Deve retornar erro quando usuário não tiver email", func(t *testing.T) {
		// arange
		usuario := Usuario{Nome: "fulano", Email: "", Senha: "sdasd"}
		expectedErrorMessage := "E-mail do usuário deve ser preenchida"

		// act
		err := ValidaUsuario(usuario)

		// assert
		assert.EqualError(t, err, expectedErrorMessage)
	})
}
