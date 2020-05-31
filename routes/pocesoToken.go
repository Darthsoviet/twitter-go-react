package routes

import (
	"errors"
	"strings"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/models"
	"github.com/Darthsoviet/twitter-go-react/util"
	"github.com/dgrijalva/jwt-go"
)

//Email valor de email usado en todos los endpoints
var Email string

//IDUsuario es el id del usuario devuelto usado en todos los modelos
var IDUsuario string

//ProcesoToken se procesa token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte(util.ObtenerEnv("JWT_CLAVE"))
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.Id
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
