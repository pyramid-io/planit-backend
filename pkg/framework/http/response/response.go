package response

import (
	"encoding/json"
	"net/http"
)

type responseInterface interface{
	Write(writer http.ResponseWriter, data interface{})
}

type jsonResponse struct{
	Status  string
	Message string
	Data    interface{}
}

func CreateJsonResponse(status string, message string, data interface{}) jsonResponse {
	return jsonResponse{
		Status: status,
		Message: message,
		Data: data,
	}
}

func (response *jsonResponse) Write(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(statusCode)
    if err := json.NewEncoder(writer).Encode(&response); err != nil {
        http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
    }
}
