package jwt

import (
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
	"github.com/Darthsoviet/twitter-go-react/util"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera el token del usuario
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte(util.ObtenerEnv("JWT_CLAVE"))
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apllidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
