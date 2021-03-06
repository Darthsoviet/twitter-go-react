package middlew

import (
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/jwt"
)

//ValidoJWT permite validar el token que nos viene en la peticion
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := jwt.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "error en el token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
