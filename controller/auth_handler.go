package controller

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

//go get ithub.com/golang-jwt/jwt
//https://jwt.io/#debugger
var jwtKey = []byte("secret")//secret para validação da assinatura

func validaToken(token string) (jwt.Claims, error) {
	tkn, err:= jwt.Parse(token, func(t *jwt.Token) (interface{}, error){
		return jwtKey, nil
	})
	if err != nil{
		return nil, err
	}
	if !tkn.Valid{
		return nil, jwt.ErrSignatureInvalid
	}
	return tkn.Claims, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Faltando authorization header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := validaToken(tokenString)

		if err != nil {
			fmt.Println("Erro ao validar jwt:")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Erro ao validar Token"))
			return
		}
		next.ServerHTTP(w, r)
		
	})
	
}

func GerarToken(w http.ResponseWriter) {
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
w.Write([]bite(tokenString))