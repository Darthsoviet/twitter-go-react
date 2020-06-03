package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
)

//ListaUsuarios lista los usuarios de la base de datos
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "debe enviar el parametro pagina como entero mayor a 0 "+err.Error(), http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, _, err := bd.LeoUsuariosTodos(jwt.IDUsuario, pag, search, typeUser)
	if err != nil {
		http.Error(w, "ocurrio un error en la base de datos "+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
