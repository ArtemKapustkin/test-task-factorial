package main

import (
	"github.com/ArtemKapustkin/test-task-factorial/internal/handler"
	"github.com/ArtemKapustkin/test-task-factorial/internal/middleware"
	"github.com/ArtemKapustkin/test-task-factorial/internal/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	calculator := service.NewFactorialCalculator()

	factorialHandler := handler.NewFactorialHandler(calculator)

	router := httprouter.New()
	router.POST("/calculate", middleware.ValidateJSON(factorialHandler.Calculate))

	err := http.ListenAndServe(":8989", router)
	if err != nil {
		log.Fatal(err)
	}
}
