//package main

//import (
//"encoding/json"
//"net/http"
//)

//func StartServer() {
//http.HandleFunc("/produtos", produtosHandler)
//http.ListenAndServe(":8080", nil)
//}

//func produtosHandler(w http.ResponseWriter, r *http.Request) {
//Produto := criaCatalogo()
//json.NewEncoder(w).Encode(Produto)
//}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/schmittalice/loja-digport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produto", produtosHandlerHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandlerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProduto(w, r)
	} else if r.Method == "POST" {
		addProduto(w, r)
	}
}

func addProduto(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	registerProduto(produto)

	w.WriteHeader(http.StatusCreated)
}

func getProduto(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		buscaPorNome := buscaPorNome(queryNome)
		json.NewEncoder(w).Encode(buscaPorNome)
	} else {
		produto := Produtos
		json.NewEncoder(w).Encode(produto)
	}
}
