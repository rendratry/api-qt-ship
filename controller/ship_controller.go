package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qt-api/helper"
	"qt-api/model/web"
)

func GetShip(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.URL.Query()
	name := params.Get("id")
	isAll := params.Get("is_all")

	filteredData := helper.FilterData(name, isAll)
	responseBody := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   filteredData,
	}

	helper.WriteToResponseBody(w, responseBody)
}
