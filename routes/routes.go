package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	ArticleRoutes(r)
	AuthRoutes(r)
}
