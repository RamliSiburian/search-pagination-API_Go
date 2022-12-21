package handlers

// import (
// 	dto "MisterAladin/dto/result"
// 	"MisterAladin/models"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type handleruser struct {
// 	UserRepository repositories.UserRepository
// }

// func HandlerUser(UserRepository repositories.UserRepository) *handleruser {
// 	return &handleruser{UserRepository}
// }

// func (h *handleruser) FindUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	user, err := h.UserRepository.FindUser()
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handleruser) GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	var user models.User
// 	user, err := h.UserRepository.GetUser(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
// 	json.NewEncoder(w).Encode(response)
// }
