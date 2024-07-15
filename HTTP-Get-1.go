package main

//importaciones
import (
"fmt"
"io/ioutil"
"log"
"net/http"
)

//Método principal
func main() {
fmt.Println("Iniciando HTTP GET Request...")
RealizarGetRequest()
}

//Método del pedido

func RealizarGetRequest() {
//crear una variable que contenga la url
const url = "172.26.10.16:400"
//crear las variables de la respuesta o el posible error
respuesta, err := http.Get(url)
//anunciar que el pedido se ha realizado
fmt.Println("Pedido Realizado")
//En caso de que haya un error parar el programa
if err != nil {
fmt.Println("Error:", err)
log.Fatal(err)
}
//Cerrar la solicitud
defer respuesta.Body.Close()
//imprimir el estado y el largo del contenido de la respuesta
fmt.Println("Codigo de estado:", respuesta.StatusCode)
fmt.Println("Longitud del codigo", respuesta.ContentLength)
//Leer el contenido de la respuesta
contenido, err := ioutil.ReadAll(respuesta.Body)
//un if para el caso de error
if err != nil {
fmt.Println("Error:", err)
log.Fatal(err)
}
//pasar a string el contenido e imprimirlo
fmt.Println(string(contenido))
}

