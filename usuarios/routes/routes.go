package routes

import (
	"github.com/go-chi/chi"
	"usuarios/controller"
	"usuarios/middleware"
)

type IRoutes interface {
	InitRoute() *chi.Mux
}

func InitRoute() *chi.Mux {
	routes := chi.NewRouter()
	routes.Use(middleware.CommonMiddleware)
	usuarioController := controller.UsuarioController{}
	rolesController := controller.RolesController{}
	routes.Get("/api/usuario", usuarioController.ListarUsarios)
	routes.Get("/api/roles", rolesController.ListarRoles)
	routes.Post("/api/usuario", usuarioController.CrearUsuario)
	routes.Delete("/api/usuario/{tipo}/{id}", usuarioController.EliminarUsuario)
	routes.Post("/api/session", usuarioController.ValidatePass)
	routes.Put("/api/usuario", usuarioController.ActualizarUsuario)
	return routes
}
