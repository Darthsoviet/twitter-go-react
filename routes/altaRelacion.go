package routes

import (
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//AltaRelacion endpoint que inserta la relacion a la base de datos
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}
	t := models.Relacion{
		UsuarioID:         jwt.IDUsuario,
		UsuarioRelacionID: ID,
	}
	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "ocurrio un error al insertar relacion "+err.Error(), http.StatusInternalServerError)

	}
	if !status {
		http.Error(w, "no se pudo insertar la relacion ", http.StatusInternalServerError)

	}
	w.WriteHeader(http.StatusCreated)

}
