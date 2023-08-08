package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
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

type calculatorDTO struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (h *FactorialHandler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody calculatorDTO

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing request body: %s", err)
		}
	}()

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	factorialA, factorialB := h.calculator.Calculate(requestBody.A, requestBody.B)

	response := map[string]uint64{
		"factorial A": factorialA,
		"factorial B": factorialB,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, "error writing http reply", http.StatusInternalServerError)
		return
	}

}
