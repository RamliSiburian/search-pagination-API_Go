package routes

import (
	"MisterAladin/handlers"
	"MisterAladin/pkg/middleware"
	"MisterAladin/pkg/mysql"
	"MisterAladin/repositories"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router) {
	articleRepositori := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepositori)

	r.HandleFunc("/articles", middleware.Auth(middleware.ArticleImage(h.CreateArticle))).Methods("POST")
	r.HandleFunc("/articles", h.FindArticle).Methods("GET")
	r.HandleFunc("/articles/{id}", h.GetArticleById).Methods("GET")
	r.HandleFunc("/articlesByuser/{user_id}", h.GetArticleByUser).Methods("GET")
	r.HandleFunc("/articlesearch/{param}", h.SearchArticle).Methods("GET")
	r.HandleFunc("/articles/{id}", h.UpdateArticle).Methods("PATCH")
	r.HandleFunc("/articles/{id}", h.DeleteArticle).Methods("DELETE")
}
