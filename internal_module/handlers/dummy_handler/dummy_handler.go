package dummy_handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pyramid.io/planit-backend/pkg/framework/application"
)

// @Summary check api work healthy with ping
// @Success 200 {string} string
// @Router /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

// @Summary check api work healthy with pong
// @Success 200 {string} string
// @Router /pong [get]
func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping!")
}

// @Summary check test logger
// @Success 200
// @Router /test-logger [get]
func TestLogger(w http.ResponseWriter, r *http.Request) {
	application.Instance.Logger.Info("this is the info")
}

// @Summary check api work healthy with pong
// @Success 200
// @Router /time [get]
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("current time is: %s", time.Now().String()))
}

// @Summary check with params functionality of router
// @Success 200
// @Router /with-params/{i}/{j} [get]
func WithParams(w http.ResponseWriter, r *http.Request) {
	i, ok := r.Context().Value("i").(string)
    if !ok {
        fmt.Println("i is missing")
        return
    }

	j, ok := r.Context().Value("j").(string)
    if !ok {
        fmt.Println("j is missing")
        return
    }

	fmt.Fprintf(w, fmt.Sprintf("i and j are: %s and %s", i, j))
}
