package main

import (
	"MisterAladin/databases"
	"MisterAladin/pkg/mysql"
	"MisterAladin/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
	}

	mysql.DatabaseInit()
	databases.Migration()

	r := mux.NewRouter()

	// routes.RouteInit(r)
	routes.RouteInit(r.PathPrefix("/api/sio").Subrouter())
	r.PathPrefix("/Uploads").Handler(http.StripPrefix("/Uploads", http.FileServer(http.Dir("./Uploads"))))

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	// var port = os.Getenv("PORT")
	var port = "5000"
	fmt.Println("server running on:" + port)

	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

}
