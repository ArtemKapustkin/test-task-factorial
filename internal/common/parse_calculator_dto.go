package common

import (
	"bytes"
	"encoding/json"
	"github.com/ArtemKapustkin/test-task-factorial/pkg"
	"io"
	"net/http"
)

type CalculatorDTO struct {
	A *int `json:"a" validate:"required,gte=0"`
	B *int `json:"b" validate:"required,gte=0"`
}

func ParseCalculatorDTO(r *http.Request) (*CalculatorDTO, *pkg.HTTPError) {
	var requestBody CalculatorDTO

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, pkg.NewHTTPError(err, http.StatusInternalServerError)
	}

	if bytes.Equal(body, nil) {
		return nil, pkg.NewHTTPError(err, http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		return nil, pkg.NewHTTPError(err, http.StatusBadRequest)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return &requestBody, nil
}
