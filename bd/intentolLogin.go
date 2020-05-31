package bd

import (
	"github.com/Darthsoviet/twitter-go-react/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin realiza el chequeo de login en la base de datos
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true

}
