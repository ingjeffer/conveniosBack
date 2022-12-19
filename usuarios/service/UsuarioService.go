package service

import (
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"golang.org/x/crypto/bcrypt"
	"usuarios/entidades"
	"usuarios/model"
	"usuarios/repository"
)

type IUsuarioService interface {
	ListarUsuarios() ([]model.Usuario, error)
	CrearUsuario(user *model.UsuarioCreate) (*model.UsuarioCreate, error)
	EliminarUsuario(id string) error
}

func ListarUsuarios() ([]model.Usuario, error) {

	var usuarioModel []model.Usuario
	usuarios := repository.ListarUsuario()
	err := dto.Map(&usuarioModel, usuarios)

	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	return usuarioModel, nil
}

func CrearUsuario(user *model.UsuarioCreate) (*model.UsuarioCreate, error) {

	var entidad entidades.Usuario

	if err := dto.Map(&entidad, user); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	encryptPass(&entidad)

	resp, err := repository.CrearUsuario(&entidad)

	if err != nil {
		return nil, err
	}

	var responseUser model.UsuarioCreate
	err = dto.Map(&responseUser, resp)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &responseUser, nil
}

func EliminarUsuario(id string, tipoId string) error {
	return repository.EliminarUsuario(id, tipoId)
}

func ActualizarUsuario(user *model.UsuarioCreate) (*model.UsuarioCreate, error) {

	var entidad entidades.Usuario
	if err := dto.Map(&entidad, user); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	encryptPass(&entidad)

	resp, err := repository.ActualizarUsuario(&entidad)

	if err != nil {
		return nil, err
	}

	var responseUser model.UsuarioCreate
	err = dto.Map(&responseUser, resp)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &responseUser, nil
}

func encryptPass(entidad *entidades.Usuario) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(entidad.Password), 6)
	entidad.Password = string(bytes)
}

func ValidatePass(email string, pass string) error {
	entidad, err := repository.GetByEmail(email)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(entidad.Password), []byte(pass))
}
