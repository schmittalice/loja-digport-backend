package model

import (
	"database/sql"
	"fmt"

	"github.com/schmittalice/loja-digport-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Nome     string `json:"nome"`
	ID       string `json:"id"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}

func ValidaUsuario(usuario Usuario) error {
	if usuario.Email == "" {
		return fmt.Errorf("Por favor, preencha o campo de E-mail")
	}
	if usuario.Senha == "" {
		return fmt.Errorf("Por favor, preencha o campo de Senha")
	}
	if usuario.Nome == "" {
		return fmt.Errorf("Por favor, preencha o campo de Nome")
	}
	return nil
}

func CriaUsuario(usuario Usuario) error {
	hash, err := hashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}
	db := db.ConectaBancoDados()

	_, err = db.Exec("INSERT INTO usuario (nome, senha, email, telefone, endereco) VALUES($1, $2, $3, $4, $5)", usuario.Nome, hash, usuario.Email, usuario.Telefone, usuario.Endereco)
	if err != nil {
		return fmt.Errorf("erro ao tentar inserir uauário no banco de dados: %w", err)
	}

	return nil
}

func hashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		return "", fmt.Errorf("erro ao tenatr gerar hash da senha: %w", err)
	}
	return string(bytes), nil
}

//func ValidaLogin(hash string, password string) error {}

func BuscaUsuarioPorEmail(email string) (*Usuario, error) {
	db := db.ConectaBancoDados()
	defer db.Close()

	var usuario Usuario
	err := db.QueryRow("SELECT id, nome, senha, email,telefone, endereco FROM usuario WHERE email = $1", email).Scan(&usuario.ID, &usuario.Nome, &usuario.Senha, &usuario.Email, &usuario.Telefone, &usuario.Endereco)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Usuário não encontrado %s", email)
		}
		return nil, err
	}

	return &usuario, nil
}
