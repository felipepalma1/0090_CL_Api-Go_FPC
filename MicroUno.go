/* Versión mínima de un micro servicio en Go. Utiliza la función http.HandleFunc para poder
enviar un mensaje por consola cada vez que es activado. Ni siquiera le maanda un mensaje al cliente.
Escucha en el puerto 9090*/

package main

import (
	"log"
	"net/http"
)

func main() {

	/*Handler básico que envía un mensaje por consola desde el lado del servidor cada vez que un
	cliente solicita el recurso raíz*/
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hola Mundo")
	})

	/*Aquí en el argumento nil es donde debe ir el Handler de la solicitud*/
	http.ListenAndServe(":9090", nil)
}
