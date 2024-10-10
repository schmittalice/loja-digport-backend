package main

import (
	"fmt"

	"github.com/schmittalice/loja-digport-backend/model"
)

func main() {

	//db.InitDB()
	//StartServer()

	//produtosFiltrados := produtosPorNome("Pacote Hungria")

	//fmt.Println(produtosFiltrados)

	prod := model.Produto{Nome: "Meia", Preco: 17.99}
	//fmt.Println("preço original do produto:", prod.Preco)
	fmt.Printf("Endereço de memória produto:%p\n", &prod)

	model.AumentaPreco(&prod)
	fmt.Println("Produto:", prod)
}

// Catalogo := criaCatalogo()
// fmt.Println("Esse é o catálogo da Loja Viaja Conosco", Catalogo)
