package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//type msg string
//func (m msg) ServeHTTP(response http.ResponseWriter, request *http.Request) {
//	fmt.Fprint(response, m)
//}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}

func main() {
	//msgHandler := msg("Hello Go")

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	//fmt.Fprint(writer, "Index Page")
	//	http.ServeFile(writer, request, "index.html")
	//	fmt.Println("Index Page Log Console")
	//})
	//http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprint(writer, "About Page")
	//})

	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
	http.Handle("/", router)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
	})

	fmt.Println("Server is listening...")
	//http.ListenAndServe(":3001", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3001", nil)
}
