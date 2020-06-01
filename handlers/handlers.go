package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Darthsoviet/twitter-go-react/middlew"
	"github.com/Darthsoviet/twitter-go-react/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto,handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoDB(routes.Registro)).Methods(http.MethodPost)
	router.HandleFunc("/login", middlew.ChequeoDB(routes.Login)).Methods(http.MethodPost)
	router.HandleFunc("/verperfil", middlew.ChequeoDB(middlew.ValidoJWT(routes.VerPerfil))).Methods(http.MethodGet)
	router.HandleFunc("/modificarPerfil", middlew.ChequeoDB(middlew.ValidoJWT(routes.ModificarPerfil))).Methods(http.MethodPut)
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidoJWT(routes.GraboTweet))).Methods(http.MethodPost)
	router.HandleFunc("/leoTweets", middlew.ChequeoDB(middlew.ValidoJWT(routes.LeoTweets))).Methods(http.MethodGet)
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
