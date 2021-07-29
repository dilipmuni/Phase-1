package main

import (
	"GOProject/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	m := httprouter.New()

	uc := controller.NewUserController(getSession())

	r.GET("/users", uc.GetAllUsers)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	m.GET("/_carts", uc.GetAllCarts)
	m.POST("/_carts", uc.CreateCart)
	m.DELETE("/_carts", uc.DeleteCart)

	m.GET("/cart", uc.GetUserCart) //needs UserID in Header
	m.PUT("/cart", uc.AddCart)
	m.DELETE("/cart", uc.DeleteItemInCart)

	// r.GET("/user/:id/payment", uc.GetPayment)
	// r.POST("/user/:id/payment", uc.PostPayment)

	// r.POST("/user/:id/order", uc.PlaceOrder)

	go http.ListenAndServe(":9012", r)
	http.ListenAndServe(":9014", m)
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
