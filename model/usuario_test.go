package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidausuario(t *testing.T) {
	t.Run("Deve retonar nulo quando usuario não tiver nome", func(t *testing.T) {
		//arrange
		usuario := Usuario{Nome: "", Email: "email", Senha: ""}
		expectedErrorMessage := "nome do usuário não pode ser vazio"

		//act
		err := ValidaUsuario(usuario)

		//assert
		assert.NoError(t, err, expectedErrorMessage)
	})

	t.Run("Deve retonar nulo quando usuario não tiver nome", func(t *testing.T) {
		//arrange
		usuario := Usuario{Nome: "", Email: "email", Senha: ""}
		expectedErrorMessage := "nome do usuário não pode ser vazio"

		//act
		err := ValidaUsuario(usuario)

		//assert
		assert.NoError(t, err, expectedErrorMessage)
	})
}
