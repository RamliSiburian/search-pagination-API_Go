package middleware

import (
	dto "MisterAladin/dto/result"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ArticleImage(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("image")
		if r.Method == "PATCH" && file == nil {
			next.ServeHTTP(w, r)
		}

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()

		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB

		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Max size in 10mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		tempFile, err := ioutil.TempFile("Uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)
		data := tempFile.Name()
		// filename := data[8:]

		ctx := context.WithValue(r.Context(), "dataFile", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
