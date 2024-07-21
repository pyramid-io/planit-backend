package dummy_handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
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

// @Summary check with params functionality of router
// @Success 200
// @Router /session-create [get]
func SessionCreateHandler(w http.ResponseWriter, r *http.Request) {
	sessionData := map[string]interface{}{
		"a": "a",
		"b": "b",
	}
	session, err := application.Instance.Session.Create(sessionData, nil)
	if err == nil {
		fmt.Fprintf(w, fmt.Sprintf(session.ID))
	}

}

// @Summary check with params functionality of router
// @Success 200
// @Router /session-get [get]
func SessionGetHandler(w http.ResponseWriter, r *http.Request) {

}

// @Summary check with params functionality of router
// @Success 200
// @Router /route-with-middleware [get]
func MiddlewareTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging in handler")
	fmt.Fprintf(w, "route with middleware handler")
}

// @Router /database-test [get]
func DatabaseTest(w http.ResponseWriter, r *http.Request) {
	connection, err := application.Instance.Database.GetConnection("api")
	if err != nil {
		fmt.Println("errror while mysql query: ", err)
	}

	collection, err := connection.Select("SELECT * FROM dummy_table")
	if err != nil {
		fmt.Println("errror while mysql query: ", err)
	}

	var dummies []dummy
	collection.Unmarshal(&dummies)

	module, err := application.Instance.GetModule("internal")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmplPath := filepath.Join(
		module.GetResourceDir(),
		"test-pages",
		"database.html",
	)

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dummies)
}

type dummy struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

// @Router / [get]
func Home(w http.ResponseWriter, r *http.Request) {
	module, err := application.Instance.GetModule("internal")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmplPath := filepath.Join(
		module.GetResourceDir(),
		"index.html",
	)

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
