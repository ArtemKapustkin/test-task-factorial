package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

var (
	validate = validator.New()
)

type calculatorDTO struct {
	A *int `json:"a" validate:"required,gte=0"`
	B *int `json:"b" validate:"required,gte=0"`
}

func ValidateJSON(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var requestBody calculatorDTO

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error occurs when read request body", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, "error parsing JSON", http.StatusBadRequest)
			return
		}

		err = validate.Struct(requestBody)
		if err != nil {
			response := map[string]string{
				"error": "Incorrect input",
			}

			responseJSON, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "error encoding JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			_, err = w.Write(responseJSON)
			if err != nil {
				http.Error(w, "error writing http reply", http.StatusInternalServerError)
				return
			}
		} else {
			if err := r.Body.Close(); err != nil {
				log.Printf("error closing request body: %s", err)
			}

			r.Body = io.NopCloser(bytes.NewBuffer(body))

			next(w, r, ps)
		}
	}

}
