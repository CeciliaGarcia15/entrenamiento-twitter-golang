package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncriptarPassword es la rutina que me permite encriptar la password

aca le pasamos un parametro pass de tipo string y la funcion va a devolver
un string y un error

*/
func EncriptarPassword(pass string) (string, error) {
	/*el costo es un numero que la funcion toma porque hace 2 elevado al costo
	para que hace eso, bueno el numero que de 2 elevado al costo es la cantidad de
	veces que va a encriptar podriamos decir va a hacer varias pasadas por el
	texto para encriptar, mientras mayor sea el costo mas seguro va a ser la contraseña
	pero al ser mas grande el numero va a pasar para encriptar mas veces y eso
	hace que tome mas tiempo hacerlo por eso tampoco podemos exagerar con el numero
	de costo porque sino la encriptacion va a tomar demasiado tiempo

	Se dice que si el usuario es un usuario normal deberia ser un costo de 6
	y si el usuario es un admin o superadmin entonces deberia ser un costo
	de 8
	*/
	costo := 8

	/*aca usamos el paquete llamado bcrypt y su metodo generate from password
	y le tenemos que pasar un slice(es como un vector pero sin tamaño definido, osea un vector
	es cuando le damos un tamaño, el slice no tiene tamaño) y a eso le tenemos que
	decir de que tipo es el slice, en este caso es un byte y a eso hay que pasasrle la
	contraseña que nos vino por parametro y despues le pasamos el costo
	*/
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	/*esto convierte en string la variable bytes y retorna eso y un err
	que si la encriptacion salio bien, err va a ser un nil*/
	return string(bytes), err
}
