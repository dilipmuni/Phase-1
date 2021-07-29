package main

import (
	"GOProject/controller"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_CreateProduct(t *testing.T) {
	uc := controller.NewUserController(getSession())

	var jsonStr = []byte(`{"pname":"test product -- 5", "pprice": 11.22, "pqty":10}`)

	create_prod_req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
	create_prod_req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product", uc.CreateProduct)
	router.ServeHTTP(rr, create_prod_req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Getting all the products
	get_allprod_req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.HandleFunc("/products", uc.GetAllProducts)
	router.ServeHTTP(rr, get_allprod_req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var m []map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	flag := 0
	for _, i := range m {
		if i["pname"] == "test product -- 5" {
			flag = 1
		}
	}
	if flag == 0 {
		t.Errorf("Product created but could not find in database")
	}

}
