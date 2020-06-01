package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/Darthsoviet/twitter-go-react/bd"
)

//ObtenerBanner endpoint para obtener el banner de un usuario mediante su id
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe de enviar parametro id", http.StatusBadRequest)
		return

	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}
	OpenFile, err := os.Open("uploads/banner/" + perfil.Banner)
	if err != nil {
		http.Error(w, "imagen no encontrada "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "imagen no encontrada "+err.Error(), http.StatusBadRequest)

	}

}
