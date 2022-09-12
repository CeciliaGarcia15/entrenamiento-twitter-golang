package bd

import (
	"context"
	"time"

	"github.com/CeciliaGarcia15/entrenamiento-twitter-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la BD para insertar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	/*aca usamos el package context y su metodo withtimeout para decirle que no quiero
	que espere una respuesta por mas de 15 segundos, y para eso le tengo que pasar el
	background que es el contexto que traigo desde la DB y tambien el tiempo que quiero
	que tenga el coxtexto ahora */
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*el defer se usa para que ejecute algo como ultima instancia(osea lo ultimo que se ehjecuta es el cancel)
	, y le pasamos el cancel para que lo ultimo que ejecute sea una cancelacion del timeout del contexto
	osea para liberar espacio o algo asi
	*/
	defer cancel()

	//aca guardo en la variable db la coneccion a la BD y le pasamos el nombre del BD
	db := MongoCN.Database("entrenamiento-twitter")
	//guardamos en col, la tabla usuarios
	col := db.Collection("usuarios")

	/*u es el parametro con el modelo de usuarios, y de ese u accedo a Password
	y ahi guardo lo que me devuelva el metodo EncriptarPassword, al que le paso
	el u.password para que lo encripte
	*/
	u.Password, _ = EncriptarPassword(u.Password)
	/*Aca en result se va a guardar la insercion de un registro a la tabla usuarios
	y le pasamos como parametro el contexto y el parametro u que tiene los datos
	*/
	result, err := col.InsertOne(ctx, u)
	/*Si hubo un error entonces retorna vacio,false y err, acordate que InsertarRegistro
	devuelve 3 parametros siempre ,por eso el return hace eso*/
	if err != nil {
		return "", false, err
	}

	/*ell id de MONGO no es un numero sino objeto raro binario,por lo tanto para poder trabajar
	con eso tenemos que hacer otras cosas , entonces decimos que objID va a guardar
	lo que traiga el result y ahi usamos el insertedID y a eso le pasamos un primitive.objectID
	y como nostros lo que queremos es un string entonces convertimos el objId a strign
	y eso lo hacemos con el .String, y como lo requiere la funcion devolvemos eso,un true y un nil para el error
	*/
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
