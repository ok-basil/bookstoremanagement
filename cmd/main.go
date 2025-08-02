package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ok-basil/bookstoremanagement/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
