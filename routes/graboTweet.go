package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/jwt"
	"github.com/Darthsoviet/twitter-go-react/models"
)

//GraboTweet ruta rest para grabar el comentario
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "formato invalido", http.StatusBadRequest)
		return
	}

	reg := models.GraboTweet{
		UserID:  jwt.IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(reg)
	if err != nil {
		http.Error(w, "error al insertar tweet a la base de datos"+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(w, "no se a logrado insertar el tweet", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
