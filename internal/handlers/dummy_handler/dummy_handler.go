package dummy_handler

import (
	"fmt"
	"net/http"

	"github.com/pyramid.io/planit-backend/pkg/framework/application"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong!")
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping!")
}

func TestLogger(w http.ResponseWriter, r *http.Request) {
	application.GetInstance().Logger.Info("this is the info")
}
