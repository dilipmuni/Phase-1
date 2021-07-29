//-------------------------------------//
//Cart Controller

package controller

import (
	"encoding/json"
	"fmt"

	//"C:/Users/Dell/Desktop/GOProject/model"
	"GOProject/model"
	"net/http"

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

func (uc UserController) GetAllCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	carts := []model.Cart{}

	if err := uc.session.DB("go-web-dev-db").C("carts").Find(nil).All(&carts); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(carts)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateCart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	cart := model.Cart{}
	json.NewDecoder(r.Body).Decode(&cart)
	cart.Id = bson.NewObjectId()

	uc.session.DB("go-web-dev-db").C("carts").Insert(cart)

	uj, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetUserCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u_id := r.Header.Get("id")

	if !bson.IsObjectIdHex(u_id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(u_id)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	cart := model.Cart{}
	if err := uc.session.DB("go-web-dev-db").C("carts").Find(bson.M{"uname": u.Name}).One(&cart); err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := r.Header.Get("id")

	if !bson.IsObjectIdHex(id) {
		fmt.Println("Inside get cart user bson error")
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("go-web-dev-db").C("carts").Remove(bson.M{"_id": oid}); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}

func (uc UserController) AddCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := r.Header.Get("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	cartproducts := model.CartProduct{}
	json.NewDecoder(r.Body).Decode(&cartproducts)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u.Name)

	cart := model.Cart{}
	if err := uc.session.DB("go-web-dev-db").C("carts").Find(bson.M{"uname": u.Name}).One(&cart); err != nil {
		fmt.Println("Error we arer in")
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	oldcartproducts := cart.CartProducts
	flag := 0

	for i, _ := range oldcartproducts {
		fmt.Println("oldcartproducts[i].ProductName", oldcartproducts[i].ProductName)
		if oldcartproducts[i].ProductName == cartproducts.ProductName {
			oldcartproducts[i].ProductQty = cartproducts.ProductQty
			flag = 1

			totalPrice := CalculateTotalPrice(oldcartproducts)
			fmt.Println("Total price:", totalPrice)
			//if err := uc.session.DB("go-web-dev-db").C("carts").Update( bson.M{"$set": bson.M{"pqty": cartproducts.ProductQty}}); err != nil {
			//if err := uc.session.DB("go-web-dev-db").C("carts").Update(bson.M{"_id": cart.Id}, bson.M{"$set": bson.M{"cartproducts": bson.M{"pqty": cartproducts.ProductQty}}}); err != nil {
			if err := uc.session.DB("go-web-dev-db").C("carts").Update(bson.M{"_id": cart.Id}, bson.M{"$set": bson.M{"cartproducts": oldcartproducts, "totalprice": totalPrice}}); err != nil {
				fmt.Println("Error we arer in")
				fmt.Println(err)
				w.WriteHeader(404)
				return
			}
		}
	}
	if flag == 0 {
		oldcartproducts = append(oldcartproducts, cartproducts)

		totalPrice := CalculateTotalPrice(oldcartproducts)
		fmt.Println("Total price:", totalPrice)

		//if err := uc.session.DB("go-web-dev-db").C("carts").Update(bson.M{"uname": u.Name}, bson.M{"cartproducts": oldcartproducts}); err != nil {
		if err := uc.session.DB("go-web-dev-db").C("carts").Update(bson.M{"_id": cart.Id}, bson.M{"$set": bson.M{"cartproducts": oldcartproducts, "totalprice": totalPrice}}); err != nil {
			fmt.Println("Error we arer in")
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
	}

	uj, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func CalculateTotalPrice(slice []model.CartProduct) float64 {
	var totalPrice float64 = 0
	for i, _ := range slice {
		totalPrice += (float64(slice[i].ProductQty) * slice[i].ProductPrice)
	}
	return totalPrice
}

func (uc UserController) DeleteItemInCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := r.Header.Get("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	cartproducts := model.CartProduct{}
	json.NewDecoder(r.Body).Decode(&cartproducts)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	cart := model.Cart{}
	if err := uc.session.DB("go-web-dev-db").C("carts").Find(bson.M{"uname": u.Name}).One(&cart); err != nil {
		fmt.Println("Error we arer in")
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	oldcartproducts := cart.CartProducts

	for i, _ := range oldcartproducts {
		fmt.Println("oldcartproducts[i].ProductName", oldcartproducts[i].ProductName)
		if oldcartproducts[i].ProductName == cartproducts.ProductName {
			oldcartproducts = remove(oldcartproducts, i)
		}
	}

	totalPrice := CalculateTotalPrice(oldcartproducts)
	fmt.Println("Total price:", totalPrice)

	//if err := uc.session.DB("go-web-dev-db").C("carts").Update(bson.M{"uname": u.Name}, bson.M{"cartproducts": oldcartproducts}); err != nil {
	if err := uc.session.DB("go-web-dev-db").C("carts").Update(bson.M{"_id": cart.Id}, bson.M{"$set": bson.M{"cartproducts": oldcartproducts, "totalprice": totalPrice}}); err != nil {
		fmt.Println("Error we arer in")
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
func remove(slice []model.CartProduct, s int) []model.CartProduct {
	return append(slice[:s], slice[s+1:]...)
}
