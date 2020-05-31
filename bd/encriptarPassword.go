package bd

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword encripta 6 veces la contraseña ddel usuario
func EncriptarPassword(pass string) (string, error) {
	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
