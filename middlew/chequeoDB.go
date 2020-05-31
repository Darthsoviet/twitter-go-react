package middlew

import (
	"net/http"

	"github.com/Darthsoviet/twitter-go-react/bd"
)

//ChequeoDB verifiica la conexion a la base de datos
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)

	}

}
