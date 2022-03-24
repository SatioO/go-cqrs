package framework

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/satioO/scheduler/scheduler/cqrs"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	app *cqrs.App
}

func NewHttpServer(app *cqrs.App) HttpServer {
	return HttpServer{app: app}
}

func HandlerFromMux(adapter HttpServer, r chi.Router) http.Handler {
	r.Post("/v1/account", adapter.OpenAccount)

	return r
}

func Run(adapter HttpServer) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", HandlerFromMux(adapter, apiRouter))

	logrus.Println("Starting HTTP server: " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), rootRouter)
}

func setMiddlewares(r *chi.Mux) {}
