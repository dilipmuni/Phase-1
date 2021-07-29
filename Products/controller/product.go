//----------------------------------------------//
//Product controller

package controller

import (
	"GOProject/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products := []model.Product{}

	if err := uc.session.DB("go-web-dev-db").C("products").Find(nil).All(&products); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	product := model.Product{}

	if err := uc.session.DB("go-web-dev-db").C("products").FindId(oid).One(&product); err != nil {
		w.WriteHeader(404)
		return
	}

	fmt.Println(product.ProductName)

	uj, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateProduct(w http.ResponseWriter, r *http.Request) {

	product := model.Product{}

	json.NewDecoder(r.Body).Decode(&product)

	product.Id = bson.NewObjectId()

	uc.session.DB("go-web-dev-db").C("products").Insert(product)

	uj, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := uc.session.DB("go-web-dev-db").C("products").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}

// html
// a-> http;/pridcts/addtocart -> ufnction -> user/api-> add toc cart

func (uc UserController) RedirectToAddCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	//http.Redirect(w, r, "http://localhost:9012/user/"+id+"/cart", http.StatusMovedPermanently)

	fmt.Println(r.Body)

	newURL := "http://localhost:9012/user/" + id + "/cart2"
	fmt.Println("Inside redirect to aadd cart 3")
	//var bdy = []byte(`{"pname":"Redmi 11","pqty":612}`)

	fmt.Println("Inside redirect to aadd cart 2")
	r.Method = "GET"
	r.URL, _ = url.Parse(newURL)
	r.RequestURI = newURL
	//.Body = ioutil.NopCloser(bytes.NewReader(bdy))
	r.Body = r.Body
	cartproducts := model.CartProduct{}
	json.NewDecoder(r.Body).Decode(&cartproducts)
	fmt.Println(cartproducts)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("Inside redirect to add cart 1")
	http.Redirect(w, r, newURL, http.StatusMovedPermanently)
}

//request and body
