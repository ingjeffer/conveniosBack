package routes

import (
	"usuarios/controller"
	"usuarios/middleware"

	"github.com/go-chi/chi"
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
	routes.Get("/api/usuario/roles", rolesController.ListarRoles)
	routes.Post("/api/usuario", usuarioController.CrearUsuario)
	routes.Delete("/api/usuario/{tipo}/{id}", usuarioController.EliminarUsuario)
	routes.Post("/api/usuario/session", usuarioController.ValidatePass)
	routes.Put("/api/usuario", usuarioController.ActualizarUsuario)
	routes.Get("/api/usuario/correo/{id}", usuarioController.ListarCorreo)
	routes.Get("/api/usuario/correo/gestor/{id}", usuarioController.ListarCorreoGestor)
	routes.Post("/api/usuario/id", usuarioController.ListarUsariosPorId)
	return routes
}
