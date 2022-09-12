package middlew

import (
	"net/http"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/bd"
)

/*ChequeoBD es el middleware que permite conocer el estado de la DB

todo handrelFunc recibe un next http.HandlerFunc
los middlewares tienen que recibir algo y devolver  lo mismo, osea el mismo t
ipo de datos porque sino no seria considerado un pasamanos,asi como recibo la
conexion del http con el response writer y el request asi como lo recibo tengo pque
pasarlo al siguiente eslabon de la cadena porque si en el body de mi request
traigo informacion yo lo voy a procesar en la ultima parte del endpoint.Entonces
asi como lo recibi lo paso por eso es que recibe un handlerFunc y por eso
envia un handlerFunc , Entonces si recibe una funcion tiene que devolver una
funcion


*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	/*Entonces retorna una funcion anonima que tiene 2 parametros llamados w y r del tipo
	http.ResponseWriter para w y un puntero a http.Request para r
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		/*si hubo un error con la conexion a la db entonces usamos
		el metodo de http llamado Error y le pasamos la variable con el
		ResponseWriter, un string sobre el error y un codigo de error, en
		este caso es el 500 y se devuelve un return lo que mata el endpoint
		y no le permite continuar*/
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		/*aca pasamos al siguiente eslabon de la cadena la informacion del response
		y el request, el next es un nombre de variable que se le paso por parametro
		y hace referencia a pasarle la info al siguiente endpoint o al siguiente eslabon
		de la cadena*/
		next.ServeHTTP(w, r)
	}
}
