package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//Registro es la funcion para crear en la bd el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recividos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "mail requerido", 400)
		return

	}
	if len(t.Password) < 6 {
		http.Error(w, "debe de especificar contraseña de almenos seis caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "ya existe usuario con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "ocurrio un error a realizar el registro de usuario"+err.Error(), 400)
		return

	}
	if status == false {
		http.Error(w, "no se pudo insertar usuario", 400)
		return

	}
	w.WriteHeader(http.StatusCreated)

}
