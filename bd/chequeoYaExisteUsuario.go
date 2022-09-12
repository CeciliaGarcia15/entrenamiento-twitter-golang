package bd

import (
	"context"
	"time"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en el BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	/*aca hacemos que el tiempo limite de respuesta sea 15 segundos*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*cuando se ejecute todo el codigo entonces se va a ejecutar el cancel() este*/
	defer cancel()

	//aca guardo en la variable db la coneccion a la BD y le pasamos el nombre del BD
	db := MongoCN.Database("entrenamiento-twitter")
	//guardamos en col, la tabla usuarios
	col := db.Collection("usuarios")

	/* usando bson.M (m es un formato) le pasamos el nombre del campo de la db
	y le colocamos como valor el email que llego por parametro */
	condicion := bson.M{"email": email}

	//crea una variable llamada resultado de tipo models.Usuario
	var resultado models.Usuario

	/*aca le decimos que en la coleccion busque UN registro no una coleccion
	sino UN registro, a eso le pasamos el contesto y la condicion , y lo decodificamos y
	guardamos en un puntero a resultado*/
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	/*desde el resultado accede a ID y eso lo convierte en un hexadecimal
	que seria como un string*/
	ID := resultado.ID.Hex()
	/*si existe un error al buscar el registro en la db entonces devuelve directamente
	el resultado, un false y el id*/
	if err != nil {
		return resultado, false, ID
	}

	/*Si no hubo error entonces devuelve el resultado,true y el id*/
	return resultado, true, ID
}
