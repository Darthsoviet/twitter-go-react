package routes

import (
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
)

//EliminoTweet endpoint encargado de borrar tweets
func EliminoTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe proporcionar un id", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(ID, jwt.IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al Borrar Id", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
