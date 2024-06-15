package jsonhandler

import (
	"net/http"

	"github.com/pyramid.io/planit-backend/pkg/framework/http/response"
)

type user struct{
	Name string
	FamilyName string
}

func JsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	response := response.CreateJsonResponse("success", "test message 2", user{Name: "mehrdad", FamilyName: "mahdian"})
	response.Write(w, 200)
}