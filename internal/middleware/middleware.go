package middleware

import (
	"github.com/ArtemKapustkin/test-task-factorial/internal/common"
	"github.com/ArtemKapustkin/test-task-factorial/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var (
	validate = validator.New()
)

func ValidateJSON(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		incorrectInput := map[string]interface{}{
			"error": "Incorrect input",
		}

		dto, httpErr := common.ParseCalculatorDTO(r)
		if httpErr != nil {
			setResponse(w, incorrectInput, http.StatusBadRequest)
			return
		}

		err := validate.Struct(dto)
		if err != nil {
			setResponse(w, incorrectInput, http.StatusBadRequest)
			return
		}

		if err := r.Body.Close(); err != nil {
			log.Printf("error closing request body: %s", err)
		}

		next(w, r, ps)
	}
}

func setResponse(w http.ResponseWriter, description map[string]interface{}, statusCode int) {
	responseJSON, httpErr := pkg.CreateJSONResponse(description)
	if httpErr != nil {
		http.Error(w, httpErr.Error(), httpErr.GetStatusCode())
	}

	httpErr = pkg.WriteJSONResponse(w, responseJSON, statusCode)
	if httpErr != nil {
		http.Error(w, httpErr.Error(), httpErr.GetStatusCode())
	}
}
