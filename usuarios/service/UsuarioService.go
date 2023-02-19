package service

import (
	"fmt"
	"usuarios/entidades"
	"usuarios/jwt"
	"usuarios/model"
	"usuarios/repository"

	"github.com/dranikpg/dto-mapper"
	"golang.org/x/crypto/bcrypt"
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

func ValidatePass(email string, pass string) (*model.Session, error) {
	entidad, err := repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(entidad.Password), []byte(pass))

	if err != nil {
		return nil, err
	}

	token, time, err := jwt.GenerarJWT(entidad)

	if err != nil {
		return nil, err
	}

	repository.GuardarSession(entidad.Id, token, time)

	return &model.Session{
		Token: token,
	}, nil
}

func ListarCorreo(role string) (*model.Usuario, error) {

	resp, err := repository.GetEmailByRole(role)

	if err != nil {
		return nil, err
	}
	fmt.Println(resp)

	return &model.Usuario{
		Id:    resp.Id,
		Email: resp.Email,
	}, nil
}

func ListarCorreoGestor(role string) (*model.Usuario, error) {

	resp, err := repository.GetUserById(role)

	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return &model.Usuario{
		Id:    resp.Id,
		Email: resp.Email,
	}, nil
}
