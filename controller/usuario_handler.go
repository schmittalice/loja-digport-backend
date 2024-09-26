package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/schmittalice/loja-digport-backend/model"
)

func CriaUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	error := model.CriaUsuario(usuario)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func BuscaUsuarioPorEmail(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")

	usuario, err := model.BuscaUsuarioPorEmail(email)
	if err != nil {
		fmt.Println("Erro ao buscar usuario:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Put(`id`)

	usuario, err := model.Usuario(id)
	if err != nil {
		fmt.Println("Erro ao editar usuario:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

//func LoginHandler(w http.ResponseWriter, r *http.Request) {
//verifica credenciais de usuario
//var usuario model.Usuario }
