package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//SubirBanner sube una imagen de usuario en el servidor
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo = "uploads/banners/" + jwt.IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al subir la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error copiando la imagen", http.StatusInternalServerError)
		return
	}
	var usuario models.Usuario
	usuario.Banner = jwt.IDUsuario + "." + extension
	status, err := bd.ModificoRegistro(usuario, jwt.IDUsuario)
	if err != nil || !status {
		http.Error(w, "error al grabar al usuario en la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
