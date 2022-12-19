package repository

import (
	"usuarios/configuration"
	"usuarios/entidades"
)

type IRolesRepository interface {
	ListarRoles() []entidades.Roles
}

func ListarRoles() []entidades.Roles {
	var roles []entidades.Roles
	configuration.Instance.Find(&roles)
	return roles
}
