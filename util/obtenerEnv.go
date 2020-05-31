package util

import (
	"os"

	"github.com/joho/godotenv"
)

func ObtenerEnv(variable string) string {
	var envVar string
	var aux string
	aux, encontrado := os.LookupEnv(variable)
	if encontrado {
		envVar = aux

	} else {
		env, _ := godotenv.Read(".env")
		envVar = env[variable]

	}
	return envVar

}
