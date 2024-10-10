package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/schmittalice/loja-digport-backend/controller"
)

func HandleRequests() {

	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	//route.HandleFunc("/produto", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controller.CriaProdutoHandler).Methods("POST")
	route.HandleFunc("/produto/{id}", controller.RemoveProdutoHandler).Methods("DELETE")
	//route.HandleFunc("/produto", controller.AtualizaProdutoHandler).methods("PUT")
	route.Handle("/produto", controller.AuthMiddleware(http.HandlerFunc(controller.BuscaProdutoPorNomeHandler))).Methods("GET")

	//Usuario
	route.HandleFunc("/usuarios", controller.CriaUsuarioHandler).Methods("POST")
	route.HandleFunc("/usuarios", controller.BuscaUsuarioPorEmail).Methods("GET")
	route.HandleFunc("/usuarios", controller.UpdateUsuario).Methods("PUT")
	route.HandleFunc("usuarios/login", controller.LoginHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)

	http.ListenAndServe(":8080", handler)
}
