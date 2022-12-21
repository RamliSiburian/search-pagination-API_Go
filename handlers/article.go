package handlers

import (
	dtoArticle "MisterAladin/dto/article"
	dto "MisterAladin/dto/result"
	"MisterAladin/models"
	"MisterAladin/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataUpload := r.Context().Value("dataFile")
	filepath := dataUpload.(string)

	request := dtoArticle.ArticleRequest{
		Title:  r.FormValue("title"),
		Body:   r.FormValue("body"),
		UserID: userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "halloCorona/articleImage"})

	if err != nil {
		fmt.Println(err.Error())
	}

	article := models.Article{
		Image:   resp.SecureURL,
		Title:   request.Title,
		Body:    request.Body,
		Created: time.Now(),
		UserID:  userId,
	}

	article, err = h.ArticleRepository.CreateArticle(article)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article, _ = h.ArticleRepository.GetArticleById(article.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: article}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerArticle) FindArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles, err := h.ArticleRepository.FindArticle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range articles {
		articles[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: articles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) GetArticleById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var article models.Article
	article, err := h.ArticleRepository.GetArticleById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article.Image = os.Getenv("PATH_FILE") + article.Image

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

func (h *handlerArticle) GetArticleByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	user_id, _ := mux.Vars(r)["user_id"]

	var article []models.Article
	article, err := h.ArticleRepository.GetArticleByUser(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range article {
		article[i].Image = os.Getenv("PATH_FILE") + p.Image
		// article[i].Image = "http://localhost:5000/Uploads/" + p.Image
	}

	w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: http.StatusOK, Data: article}
	json.NewEncoder(w).Encode(article)
}

func (h *handlerArticle) SearchArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	param, _ := mux.Vars(r)["param"]

	var article []models.Article
	article, err := h.ArticleRepository.SearchArticle(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

func (h *handlerArticle) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dtoArticle.ArticleRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.ArticleRepository.GetArticleById(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		user.Title = request.Title
	}

	if request.Body != "" {
		user.Body = request.Body
	}

	data, err := h.ArticleRepository.UpdateArticle(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	article, err := h.ArticleRepository.GetArticleById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteArticle, err := h.ArticleRepository.DeleteArticle(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteArticle}
	json.NewEncoder(w).Encode(response)
}
