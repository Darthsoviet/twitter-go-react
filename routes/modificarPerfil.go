package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//ModificarPerfil endpoint para modificar  implementa el tipo de HandleFunc
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "datos incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar modificar el registro, reintente nuevamente"+err.Error(), http.StatusBadRequest)
		return

	}
	if status == false {
		http.Error(w, "no se a logrado modificar el registro del usuario", http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusCreated)

}
