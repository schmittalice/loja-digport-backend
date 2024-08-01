package routes

import
"net/http"

"Github.com/gorilla/mux"
"https://github.com/schmittalice/loja-digiport-backend"

func HandleRequests() route := mux.NewRouter()

route.handleFunc("/produtos, controller.BuscaProdutosHandler).Methods("GET")
route.handleFunc("/produtos, controller.BuscaProdutoPorNomeHandler).Methods("GET")
route.handleFunc("/produtos, controller.CriaProdutosHandler).Methods("GET")

http.ListenAndServe(":8080", route)
