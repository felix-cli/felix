package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/update/update_me/internal/config"
	"go.uber.org/zap"
)

// Handler is the struct to track dependencies of the Vendor Portal Service routes.
type Handler struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

// New creates the route handler with the dependencies in the parameters.
func New(ll *zap.SugaredLogger, cfg *config.Config) *mux.Router {
	h := &Handler{
		Log:    ll,
		Config: cfg,
	}

	r := mux.NewRouter()

	r.HandleFunc("/hello_world", h.HelloWorld).Methods("GET")

	return r
}

// HelloWorld is a sample endpoint to show that your sevice is working
func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Hello World`))
}
