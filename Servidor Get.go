package main
import(
	"fmt"
	"net/http"
	"html/template"
)
func index(w http.ResponseWriter, r http.Request){
	template, err := template.ParseFiles("templates/index.html")
	if err != nil{
		fmt.Fprintf("Página no encontrada")
	} else{
		template.Execute(w, nil)
	}
}
func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":400", nil)
	fmt.Println("El servidor está a la escucha en el puerto 400")