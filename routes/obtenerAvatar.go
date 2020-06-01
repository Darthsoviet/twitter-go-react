package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/Darthsoviet/twitter-go-react/bd"
)

//ObtenerAvatar endpoint para obtener el avatar de un usuario mediante su id
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
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
	OpenFile, err := os.Open("uploads/avatar/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "imagen no encontrada "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "imagen no encontrada "+err.Error(), http.StatusBadRequest)

	}

}
