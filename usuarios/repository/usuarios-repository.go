package repository

import (
	"fmt"
	"usuarios/configuration"
	"usuarios/entidades"
)

var usuarioDuplicadoError usuarioError

type usuarioError struct {
	msg string
}

func (user usuarioError) Error() string {
	return user.msg
}

func ListarUsuario() []entidades.Usuario {
	var usuarios []entidades.Usuario
	configuration.Instance.Model(&entidades.Usuario{}).Preload("Roles").Find(&usuarios)
	return usuarios
}

func ListarUsuariosPorIDs(userIDs []string) ([]entidades.Usuario, error) {
	var usuarios []entidades.Usuario
	result := configuration.Instance.Model(&entidades.Usuario{}).
		Preload("Roles").
		Where("id IN (?)", userIDs).
		Find(&usuarios)
	if result.Error != nil {
		return usuarios, result.Error
	}
	return usuarios, nil
}
func CrearUsuario(usuario *entidades.Usuario) (*entidades.Usuario, error) {

	result := configuration.Instance.Create(&usuario)
	if err := result.Error; err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Error creando usuario")
	}
	return usuario, nil
}

func ActualizarUsuario(usuario *entidades.Usuario) (*entidades.Usuario, error) {
	result := configuration.Instance.Updates(&usuario)
	if err := result.Error; err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Error actualizando usuario")
	}
	return usuario, nil
}

func EliminarUsuario(id string, tipoId string) error {
	result := configuration.Instance.Select("Roles").Delete(&entidades.Usuario{Id: id, TipoId: tipoId})

	if err := result.Error; err != nil {
		return fmt.Errorf("Error eliminando usuario")
	}
	return nil
}

func GetByEmail(email string) (*entidades.Usuario, error) {
	var usuario entidades.Usuario
	result := configuration.Instance.Preload("Roles").Where("Email = ?", email).Find(&usuario)

	if err := result.Error; err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Error buscando por email")
	}
	return &usuario, nil
}

func GetEmailByRole(role string) (*entidades.Usuario, error) {
	var usuario entidades.Usuario

	result := configuration.Instance.Joins("JOIN roles ON usuarios.role_id = roles.id").Where("roles.nombre = ?", role).Preload("Roles").Find(&usuario)

	if err := result.Error; err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Error buscando por email")
	}

	return &usuario, nil
}

func GetUserById(id string) (*entidades.Usuario, error) {
	var usuario entidades.Usuario
	fmt.Println(id)
	result := configuration.Instance.Where("id = ?", id).Find(&usuario)

	if err := result.Error; err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Error buscando por id")
	}

	return &usuario, nil
}
