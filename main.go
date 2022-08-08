package main

import (
	"log"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/bd"
	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
