package router

import (
	"net/http"

	openapimiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/hurbcom/jobs-manager-gke/server/handler"
	"github.com/hurbcom/jobs-manager-gke/server/router/routes"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	routes.GenerateRoutes(r)
	SetupMinimalHandlers(r)
	SetupSwaggerHandlers(r)
	return r

}

func SetupMinimalHandlers(r *mux.Router) {
	r.HandleFunc("/healthcheck", handler.HealthcheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/readiness", handler.ReadinessHandler).Methods(http.MethodGet)
}

func SetupSwaggerHandlers(r *mux.Router) {
	opts := openapimiddleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := openapimiddleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./swagger")))
}
