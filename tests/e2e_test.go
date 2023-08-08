package tests

import (
	"bytes"
	"github.com/ArtemKapustkin/test-task-factorial/internal/handler"
	"github.com/ArtemKapustkin/test-task-factorial/internal/middleware"
	"github.com/ArtemKapustkin/test-task-factorial/internal/service"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"testing"
)

func TestFactorialHandler(t *testing.T) {
	calculator := service.NewFactorialCalculator()

	factorialHandler := handler.NewFactorialHandler(calculator)

	router := httprouter.New()
	router.POST("/calculate", middleware.ValidateJSON(factorialHandler.Calculate))

	go func() {
		err := http.ListenAndServe(":8989", router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	tests := []struct {
		name               string
		expectedStatusCode int
		body               string
	}{
		{
			name:               "Success Calculation #1",
			expectedStatusCode: http.StatusOK,
			body:               `{"a":0,"b":1}`,
		},
		{
			name:               "Success Calculation #2",
			expectedStatusCode: http.StatusOK,
			body:               `{"a":4,"b":20}`,
		},
		{
			name:               "Success Calculation #3",
			expectedStatusCode: http.StatusOK,
			body:               `{"a":6,"b":9}`,
		},
		{
			name:               "Bad Request Error #1",
			expectedStatusCode: http.StatusBadRequest,
			body:               `{"a":15,"b":-6}`,
		},
		{
			name:               "Bad Request Error #2",
			expectedStatusCode: http.StatusBadRequest,
			body:               `{"a":-10,"b":7}`,
		},
		{
			name:               "Bad Request Error #3",
			expectedStatusCode: http.StatusBadRequest,
			body:               ``,
		},
		{
			name:               "Bad Request Error #4",
			expectedStatusCode: http.StatusBadRequest,
			body:               `{}`,
		},
	}

	client := &http.Client{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodPost, "http://localhost:8989/calculate", bytes.NewBufferString(test.body))
			require.NoError(t, err)

			request.Header.Set("Content-Type", "application/json")

			response, err := client.Do(request)
			require.NoError(t, err)

			err = response.Body.Close()
			require.NoError(t, err)

			require.Equal(t, test.expectedStatusCode, response.StatusCode)
		})
	}
}
