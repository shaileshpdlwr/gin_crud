package main
import (
	"encoding/json"
	"net/http"
"fmt"
"github.com/gorilla/mux"  // for http router
)

func main(){
	router := mux.NewRouter() //To create new router

	router.HandleFunc("/books",func(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode("Hello World") 	
}) //url and output
	fmt.Println("Hello World")
	http.ListenAndServe(":4000",router)
}
