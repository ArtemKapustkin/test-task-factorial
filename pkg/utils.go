package pkg

import (
	"encoding/json"
	"net/http"
)

func CreateJSONResponse(response map[string]interface{}) ([]byte, *HTTPError) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return nil, NewHTTPError(err, http.StatusInternalServerError)
	}

	return responseJSON, nil
}

func WriteJSONResponse(w http.ResponseWriter, responseJSON []byte, statusCode int) *HTTPError {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err := w.Write(responseJSON)
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError)
	}

	return nil
}
