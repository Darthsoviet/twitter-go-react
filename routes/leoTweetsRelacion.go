package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
)

//LeoTweetsRelacion endpoint para leer los tweets de los seguidoresS
func LeoTweetsRelacion(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "pagina debe de ser mayor a 0", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}
	respuesta, correcto, err := bd.LeoTweetSeguidores(jwt.IDUsuario, pagina)

	if !correcto {
		http.Error(w, "error al leer los tweets ", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "error con la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respuesta)
}
