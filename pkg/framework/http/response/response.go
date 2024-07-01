package response

import (
	"encoding/json"
	"net/http"
)

type ResponseInterface interface{
	Write(writer http.ResponseWriter, data interface{})
}

type JsonResponse struct{
	Status  string
	Message string
	Data    interface{}
}

func CreateJsonResponse(status string, message string, data interface{}) JsonResponse {
	return JsonResponse{
		Status: status,
		Message: message,
		Data: data,
	}
}

func (response *JsonResponse) Write(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(statusCode)
    if err := json.NewEncoder(writer).Encode(&response); err != nil {
        http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
    }
}
