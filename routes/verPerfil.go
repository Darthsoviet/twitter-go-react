package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
)

//VerPerfil permite extraer los valores de perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "enviar el parametro id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "ocurrio error al intentar buscar el registro"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)

}
