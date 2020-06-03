package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//ConsultaRelacion endpoint para consultar relaciones entre usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	t := models.Relacion{
		UsuarioID:         jwt.IDUsuario,
		UsuarioRelacionID: ID,
	}
	status, err := bd.ConsultoRelacion(t)
	if err != nil {
		http.Error(w, "error en la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}
	st := models.RespuestaConsultaRelacion{
		Status: status,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(st)
	w.WriteHeader(http.StatusOK)

}
