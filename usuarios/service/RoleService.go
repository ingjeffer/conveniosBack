package service

import (
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"usuarios/entidades"
	"usuarios/model"
	"usuarios/repository"
)

type IRoleService interface {
	ListarRoles() ([]entidades.Roles, error)
}

func ListarRoles() ([]model.Roles, error) {

	var modelRoles []model.Roles
	roles := repository.ListarRoles()
	err := dto.Map(&modelRoles, roles)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return modelRoles, nil
}
