package routes

import (
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//BajaRelacion se encarga de comunicarse con la bd y borrar un registro de relacion
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	t := models.Relacion{
		UsuarioID:         jwt.IDUsuario,
		UsuarioRelacionID: ID,
	}
	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "ocurrio un error en la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "ocurrio un error inesperado", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
