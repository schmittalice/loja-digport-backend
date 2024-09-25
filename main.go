package main

import "github.com/schmittalice/loja-digport-backend/db"

func main() {

	db.InitDB()
	StartServer()

	//produtosFiltrados := produtosPorNome("Pacote Hungria")

	//fmt.Println(produtosFiltrados)
}

// Catalogo := criaCatalogo()
// fmt.Println("Esse é o catálogo da Loja Viaja Conosco", Catalogo)
