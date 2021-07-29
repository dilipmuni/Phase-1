package main

import (
	"GOProject/controller"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_GetAllProduct(t *testing.T) {

	uc := controller.NewUserController(getSession())

	get_allprod_req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	//get_prod_req,err := http.NewRequest("GET","/product",)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products", uc.GetAllProducts)
	router.ServeHTTP(rr, get_allprod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var m []map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	//fmt.Println(m[0])
}

func Test_GetProduct(t *testing.T) {

	uc := controller.NewUserController(getSession())

	get_prod_req, err := http.NewRequest("GET", "/product", nil)
	get_prod_req.Header.Set("Content-Type", "application/json")
	get_prod_req.Header.Set("id", "60f521c9a1875b06b86c620e")
	if err != nil {
		t.Fatal(err)
	}

	//get_prod_req,err := http.NewRequest("GET","/product",)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product", uc.GetProduct)
	router.ServeHTTP(rr, get_prod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	fmt.Println(m)
}

func Test_DeleteProduct(t *testing.T) {

	uc := controller.NewUserController(getSession())

	delete_prod_req, err := http.NewRequest("DELETE", "/product", nil)
	delete_prod_req.Header.Set("Content-Type", "application/json")
	delete_prod_req.Header.Set("id", "60fe815aa1875b19900b982d")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product", uc.DeleteProduct)
	router.ServeHTTP(rr, delete_prod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
