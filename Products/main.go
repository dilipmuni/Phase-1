package main

import (
	"GOProject/controller"
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {

	r := mux.NewRouter()
	uc := controller.NewUserController(getSession())
	//fmt.Println(uc.session)

	r.HandleFunc("/products", uc.GetAllProducts).Methods("GET")
	r.HandleFunc("/product", uc.GetProduct).Methods("GET") //needs userid in Header
	r.HandleFunc("/product", uc.CreateProduct).Methods("POST")
	r.HandleFunc("/product", uc.DeleteProduct).Methods("DELETE") //needs userid in Header

	log.Fatal(http.ListenAndServe(":9013", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
func fun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
