package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qt-api/controller"
	"qt-api/helper"
)

func main() {
	router := httprouter.New()

	router.GET("/api/ship", controller.GetShip)

	err := http.ListenAndServe(":8080", router)
	helper.PanicIfError(err)
}
