package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*manejadores seteo mi puerto, el handler y pogo a escuchar el servidor*/
func Manejadores() {
	/*guarda en router una nueva ruta que eso lo hacemos usando del paquete mux la funcion newRouter*/
	router := mux.NewRouter()

	/*Guarda en la variable PORT el puerto si es que tenemos uno definido, si no lo tenemos
	entonces guarda null, si eso sucede entonces le colocamos nosotros el puerto*/
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	/*allowall da todos los permisos para todo el que se conecte, no pone restricciones*/
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
