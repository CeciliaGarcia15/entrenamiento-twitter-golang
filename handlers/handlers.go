package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/middlew"
	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Que hace mux: captura el http y le da un manejo al resposewriter y al request que viene de la API
y va a poder procesar, va a ver en el body del llamada hay info si en el header hay
info y va a poder mandar una respuesta al navegador ntonces cuando llamemmos una ruta de
la api, la api nos va a devolverun status como por ej si se creo un registro correctamente, si
me pude loguear, me va a devolver un token y va a devolver info a la vez que recibe info
y todo eso lo maneja la rutina mux. y ese es el router que se crea en el manejador


que son los cors: Son los permisos que le doy a la api para que sea accesible desde cualquier lugar
*/

/*manejadores seteo mi puerto, el handler y pongo a escuchar el servidor*/
func Manejadores() {
	/*guarda en router una nueva ruta que eso lo hacemos usando del paquete mux la funcion newRouter*/
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

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
