package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la BD*/

var MongoCN = ConectarBD()

/*clientOptions  es donde tenemos guardado la conexion a mongo*/
var clientOptions = options.Client().ApplyURI("mongodb+srv://cecilia_garcia:alhz4twjI9TjoAI6@entrenamiento-twitter-g.4ygwyc6.mongodb.net/?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la BD y devuelve un cliente de mongo*/
func ConectarBD() *mongo.Client {
	/*esto conecta a la db, se conecta a la url que le pasamos en la variable clientOptions
	y le pasamos el context.TODO que significa que el contexto no tiene ninguna restriccion
	de tiempo o cosas asi para la conexion
	CONTEXT =  espacio de memoriavpara compartir cosas, nos sirven para comunicar
	informacion entre ejecuciones . sirven(context) porque podemos configurarles
	tiempos de espera o timeouts y si la db no responde dentro de ese tiempo entonces
	no pasa nada y la API no se cuelga, porque si se colgara entonces todas las peticiones
	siguientes tambien se cuelgan
	Esta funcion nos va a devolver dos resultados un client que es el que va a guardar
	la conexion a la db y un err que justo por eso abajo pregunta que si err tiene algo
	entonces muestro por consola el error y devuelvo client pero porque? bueno porque cuando
	creamos la funcion le dijimos que el retorno o lo que iba a devolver la funcion
	es un mongo.Client osea un objeto de mongo por eso siempre aunque pase algun error
	hay que retornar el client aunque este vacio poruqe ocurrio un error, igual vacio y
	todo tenes que devolverlo
	*/
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	/*aca hace un chequeo del ping , es lo mismo casi que lo que esta en la funcion de abajo
	solo que esta retorna un client y ademas muestra el error por consola
	aca el err no lleva el := porque el := se utiliza cuando la variable se crea por primera vez en la rutina
	en este caso err ya se creo arriba junto con client cuando es := crea la variable y
	le asigna un valor todo al mismo tiempo  pero eso se hace solo cuando es una variable nueva	cuando es para el uso de la variable entonces es solo con el =.
	Aca importante, del error que muestro con el log hay que mostrar la funcion Error porque o
	si no te puede dar otro error al mostrar*/
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	/*Si todo funciono entonces mostramos por consola el mensaje Conexion exitosa con la db
	y retorna un client*/
	log.Println("Conexion Exitosa con la BD")
	return client
}

/*ChequeoConnection es el Ping a la BD
hace un ping para saber si la db esta arriba, se debe hacer antes de realizar
ninguna instruccion para asegurarnos de que la db esta arriba, disponible porque
si no nos va a dar un monton de errores solo porque estamos tratando, luego le
pasamos el contexto (te explique mas o menos lo que era en la funcion de arriba)
y luego un nil que es un parametro que te pide la propia funcion Ping, entonces
lo que hace es que guarda en la variable err el ping , si todo sale bien significa
que la db funciona perfecto, por el contrario si hubo un error entonces entraria en
el if y retornaria 0 sino retorna 1, osea si la db funciona retorna 1 , porque???
porque en la funcion colocamos que iba a devolver un entero(int) pero tranquilamente
podriamos cambiar eso y colocar que devuelva un booleano
Si viste que en la funcion de arriba es la variable client la que llama a la funcion
.Ping() ,bueno aca no porque aca afuera de esa funcion la forma de conectarme a la base de
datos es mediante la variable o el objeto MongoCN, esa variable es la que tiene todo lo
referente a la conexion
*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
