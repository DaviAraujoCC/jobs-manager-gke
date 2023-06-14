package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hurbcom/jobs-manager-gke/internal/middleware"
	"github.com/hurbcom/jobs-manager-gke/server/handler"
)

type Routes struct {
	// Path is the path of the route
	Path string
	// Methods is the HTTP methods of the route
	Methods []string
	// Handler is the handler function of the route
	Handler func(http.ResponseWriter, *http.Request)
	// AuthRequired is a flag that indicates if the route requires authentication
	AuthRequired bool
}

var cronJobRoutes = []Routes{
	{
		Path:         "/cronjobs",
		Methods:      []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions},
		Handler:      handler.CronJobsHandler,
		AuthRequired: true,
	},
	{
		Path:         "/cronjobs/{name}",
		Methods:      []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions},
		Handler:      handler.CronJobsHandler,
		AuthRequired: true,
	},
}

var jobRoutes = []Routes{
	{
		Path:         "/jobs",
		Methods:      []string{http.MethodGet, http.MethodOptions},
		Handler:      handler.JobsHandler,
		AuthRequired: true,
	},
	{
		Path:         "/jobs/{name}",
		Methods:      []string{http.MethodGet, http.MethodOptions},
		Handler:      handler.JobsHandler,
		AuthRequired: true,
	},
}

func GenerateRoutes(r *mux.Router) {

	rrs := []Routes{}

	rrs = append(rrs, cronJobRoutes...)
	rrs = append(rrs, jobRoutes...)

	r.Handle("/", http.RedirectHandler("/docs", http.StatusMovedPermanently)).Methods(http.MethodGet, http.MethodOptions)

	for _, route := range rrs {
		if route.AuthRequired {
			r.HandleFunc(route.Path, middleware.CORS(middleware.Authenticate(route.Handler))).Methods(route.Methods...)
		} else {
			r.HandleFunc(route.Path, middleware.CORS(route.Handler)).Methods(route.Methods...)
		}
	}

}
