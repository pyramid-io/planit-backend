package json_handler

import (
	"net/http"

	"github.com/pyramid.io/planit-backend/pkg/framework/http/response"
)

type User struct{
	Name string
	FamilyName string
}

func JsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	response := response.CreateJsonResponse("success", "test message 2", User{Name: "mehrdad", FamilyName: "mahdian"})
	response.Write(w, 200)
}