package main

import (
	"log"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/bd"
	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/handlers"
)

/* Si falla la conexion entonces te manda un log que dice sin conexion a la db y me retorna
y despues llama a Manejadores dentro de la carpeta handlers*/

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
