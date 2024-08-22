package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schmittalice/loja-digport-backend/controller"
)

func HandleRequests() {

	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controller.CriaProdutoHandler).Methods("POST")
	//route.HandleFunc("produto", controller.RemoveProdutosHandler, r *http.Request) ("DELETE")
	//route.HandleFunc("/produto", controllers.AtualizaProdutoHandler).methods("PUT")

	http.ListenAndServe(":8080", route)
}
