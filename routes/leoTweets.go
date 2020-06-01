package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Darthsoviet/twitter-go-react/bd"
)

//LeoTweets leo los tweets
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe enviar el parametro pagina", http.StatusBadRequest)
		return

	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "pagina debe se mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweet(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
	log.Print(respuesta)

}
