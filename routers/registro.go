package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/bd"
	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/models"
)

/*Registro es la funcion para crear en la DB el registro de usuario*/

func Registro(w http.ResponseWriter, r *http.Request) {
	// guardamos en t el modelo de usuario
	var t models.Usuario
	/**/
	err := json.NewDecoder(r.Body).Decode(&t)
	/*si hay un error devuelve, el responseWriter, un string concatenado con el error
	que salio y por ultimo el codigo  del error 400*/
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	/*Si la cantidad de letras que tiene el email que nos llego es igual a 0
	significa que el campo esta vacio, asi que envia el error y retorna vacio
	para que se corte directamente la ejecucion */
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	/*Si la cantidad de letras que tiene la contraseña que nos llego es menor a 6
	entonces no cumple con la condicion y envia el error y retorna vacio
	para que se corte directamente la ejecucion */
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	/*aca devuelve tres valores, entonces como solo nos importa el encontrado
	los otros dos no les ponemos nombre y solo colocamos guiones bajos
	y ahi llamamos al metodo ChequeoYaExisteusuario del paquete bd y le pasamos el
	email , si encontrado es verdadero, significa que ese usuario ya existia en la
	db y por eso se coloca el error y no retorna nada para que corte la ejecucion de
	este metodo
	*/
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	/*aca tambien devuelve 3 valores, pero solo nos interesan 2 de esos valores
	que son el status y el err , aca se usa el paquete  bd y el metodo InsertoRegistro
	y le pasamos todos los datos que estan guardados en la variable t*/
	_, status, err := bd.InsertoRegistro(t)
	/*Si hay un error entonces usa ek http.Error y se le pasa el response,el mensaje
	concatenado con el error y por ultimo el codigo http correspondiente y se corta la
	ejecucion usando el return vacio*/
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	/*Si el status es falso entonces significa que algo salio mal y el registro no se guardo
	como siempre usas el http.error y retornas vacio para que corte la ejecucion
	*/
	if status == false {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario", 400)
		return
	}
	/*Si todo salio bien y no ocurrio ningun error entonces usamos el metodo WriteHeader
	que esta definido en el ResponseWriter y se le pasa el metodo de http llamado
	StatusCreated*/

	w.WriteHeader(http.StatusCreated)
}
