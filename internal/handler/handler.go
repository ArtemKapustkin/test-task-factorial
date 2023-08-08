package handler

import (
	"github.com/ArtemKapustkin/test-task-factorial/internal/common"
	"github.com/ArtemKapustkin/test-task-factorial/pkg"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Calculator interface {
	Calculate(a, b int) (uint64, uint64)
}

type FactorialHandler struct {
	calculator Calculator
}

func NewFactorialHandler(calculator Calculator) *FactorialHandler {
	return &FactorialHandler{
		calculator: calculator,
	}
}

func (h *FactorialHandler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dto, httpErr := common.ParseCalculatorDTO(r)
	if httpErr != nil {
		http.Error(w, httpErr.Error(), httpErr.GetStatusCode())
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing request body: %s", err)
		}
	}()

	factorialA, factorialB := h.calculator.Calculate(*dto.A, *dto.B)

	response := map[string]interface{}{
		"factorial A": factorialA,
		"factorial B": factorialB,
	}

	responseJSON, httpErr := pkg.CreateJSONResponse(response)
	if httpErr != nil {
		http.Error(w, httpErr.Error(), httpErr.GetStatusCode())
	}

	httpErr = pkg.WriteJSONResponse(w, responseJSON, http.StatusOK)
	if httpErr != nil {
		http.Error(w, httpErr.Error(), httpErr.GetStatusCode())
	}
}
