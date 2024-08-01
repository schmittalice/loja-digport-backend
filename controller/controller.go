package controller

import (
	"encoding/json"
	"https://github.com/schmittalice/loja-digport-backend/model"
	"net/http"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {

}
