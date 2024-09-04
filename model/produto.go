package model

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/schmittalice/loja-digport-backend/db"
)

type Produto struct {
	Nome                string  `json:"nome"`
	Descricao           string  `json:"descricao"`
	Categoria           string  `json:"categoria"`
	ID                  string  `json:"id"`
	Preco               float64 `json:"preco"`
	Quantidade          int     `json:"quantidade"`
	Imagem              string  `json:"imagem"`
	QuantidadeDeDias    int     `json:"quantidadeDeDias"`
	QuantidadeEmEstoque int     `json:"quantidadeEmEstoque"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
		// return nil, errors.New("erro ao buscar produtos")
	}

	p := Produto{}
	produtos := []Produto{}

	for resultado.Next() {

		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()

	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto nao encontrado %s\n", nome)

	} else if err != nil {
		panic(err.Error())
	}

	var p Produto
	p.ID = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Imagem = imagem
	p.QuantidadeEmEstoque = quantidade

	defer db.Close() //sempre colocar para não ficar conexao aberta
	return p
}

func produtoCadastrado(nomeProduto string) bool {

	prod := BuscaProdutoPorNome(nomeProduto)

	return prod.Nome == nomeProduto
}

func CriaProduto(prod Produto) error {
	fmt.Println("Nome do produto que quero cadastrar: ", prod.Nome)
	if produtoCadastrado(prod.Nome) {
		fmt.Printf("Produto já cadastrado: %s\n", prod.Nome)
		return fmt.Errorf("Produto já cacastrado")
	}

	db := db.ConectaBancoDados()
	id := uuid.NewString()
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	strInsert := "INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)"

	result, err := db.Exec(strInsert, id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade))
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())

	}
	fmt.Printf("Produto %s criado com sucesso (%d row affected)\n", id, rowsAffected)

	defer db.Close()

	return nil
}

func RemoveProduto(id string) error {
	db := db.ConectaBancoDados()

	resultado, err := db.Exec("DELETE FROM PRODUTOS WHERE id = $1", id)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao tentar excluir produto: %s", err.Error())
		return fmt.Errorf("Ocorreu um erro ao tentar excluir o produto: %w", err)
	}

	linesAffected, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("%d linhas afetadas\n", linesAffected)
	fmt.Println("Produto excluído")

	defer db.Close()
	return nil
}
