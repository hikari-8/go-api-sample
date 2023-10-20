package web

import (
	"log"
	"net/http"

	"github.com/mahiro72/go-api-sample/web/controller"

	"github.com/go-chi/chi/v5"
)

func Serve() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.CreateUser)
		r.Get("/{userID}",controller.GetUser)
	})

	if err:=http.ListenAndServe(":8080",r);err!=nil{
		log.Fatal(err)
	}
}