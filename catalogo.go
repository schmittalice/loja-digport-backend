package main

import (
	"github.com/schmittalice/loja-digport-backend/model"
)

func criaCatalogo() []model.Produto {
	Produtos := []model.Produto{
		{
			ID:                  "1",
			Nome:                "Pacote Espanha",
			Valor:               22.000,
			Descricao:           "Pacote de viagem para Espanha em 2025",
			Imagem:              "Barcelona.jpg",
			QuantidadeDeDias:    15,
			QuantidadeEmEstoque: 20,
		},
		{

			ID:                  "2",
			Nome:                "Pacote Grecia",
			Valor:               20.000,
			Descricao:           "Pacote de viagem para Grecia em 2025",
			Imagem:              "Atenas.jpg",
			QuantidadeDeDias:    15,
			QuantidadeEmEstoque: 20,
		},
		{
			ID:                  "3",
			Nome:                "Pacote Hungria",
			Valor:               18.000,
			Descricao:           "Pacote de viagem para Hungria em 2025",
			Imagem:              "Budapeste.jpg",
			QuantidadeDeDias:    15,
			QuantidadeEmEstoque: 20,
		},
	}
	return Produtos
}

func produtosPorNome(nome string) []model.Produto {
	catalogo := []model.Produto{}
	Produtos := criaCatalogo()

	for i := range Produtos {
		produtoBuscado := Produtos[i]
		if produtoBuscado.Nome == nome {
			catalogo = append(catalogo, produtoBuscado)
		}
	}
	return catalogo
}
