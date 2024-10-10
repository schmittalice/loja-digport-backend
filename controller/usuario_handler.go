package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/schmittalice/loja-digport-backend/model"
)

var jwtKey = []byte("secret")

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

	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	user, err := model.UpdateUsuario(usuario)
	if err != nil {
		fmt.Println("Erro ao editar usuario:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

//func LoginHandler(w http.ResponseWriter, r *http.Request) {
//verifica credenciais de usuario
//	var usuario model.Usuario
//}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	//senha := usuario.Senha
	username := usuario.Email
	user, error := model.BuscaUsuarioPorEmail(username)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	hash := user.Senha

	error = model.ValidaLogin(hash, senhatxt)

	if error == nil {
		GerarToken(w)
	} else {
		w.WriteHeader(http.StatusUnauthorized)

	}
}

func gerarToken(w http.ResponseWriter) {
	dataExpiracao := time.Now().Add(60 * time.Minute)
	standardToken := jwt.StandardClaims{
		ExpiresAt: dataExpiracao.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardToken)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println("Erro ao validar jwt:")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(tokenString))
	return
}
