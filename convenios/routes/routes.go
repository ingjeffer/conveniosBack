package routes

import (
	"convenios/controller"
	"convenios/middleware"

	"github.com/go-chi/chi"
)

type IRoutes interface {
	InitRoute() *chi.Mux
}

func InitRoute() *chi.Mux {
	routes := chi.NewRouter()
	convenioController := controller.ConveniosController{}
	routes.Use(middleware.CommonMiddleware)
	routes.Post("/api/convenio", convenioController.CrearConvenio)
	routes.Get("/api/convenio", convenioController.GetConvenios)
	routes.Get("/api/convenio/{id}", convenioController.GetConvenio)
	routes.Put("/api/convenio", convenioController.ActualizarConvenio)
	routes.Get("/api/convenio/pdf/{id}", convenioController.GenerarPDFConvenio)
	routes.Post("/api/convenio/cambiarEstado/{id}", convenioController.CambiarEstado)

	return routes
}
