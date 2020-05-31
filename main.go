package main

import (
	"log"

	"github.com/Darthsoviet/twitter-go-react/bd"
	"github.com/Darthsoviet/twitter-go-react/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a BD")
		return
	}
	handlers.Manejadores()
}
