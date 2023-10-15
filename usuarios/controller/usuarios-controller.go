package controller

import (
	"encoding/json"
	"net/http"
	"usuarios/entidades"
	"usuarios/model"
	"usuarios/service"

	"github.com/go-chi/chi"
)

type UsuarioController struct {
}

func (controller *UsuarioController) ListarUsariosPorId(w http.ResponseWriter, r *http.Request) {

	type Datos struct {
		Users []string `json:"users"`
	}

	var data Datos

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Error al decodificar el JSON del cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	usuarios, err := service.ListarUsuariosPorId(data.Users)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func (controller *UsuarioController) ListarUsarios(w http.ResponseWriter, r *http.Request) {

	usuarios, err := service.ListarUsuarios()

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func (controller *UsuarioController) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var user model.UsuarioCreate
	json.NewDecoder(r.Body).Decode(&user)

	userSave, err := service.CrearUsuario(&user)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userSave)

}

func (controller *UsuarioController) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var user model.UsuarioCreate
	json.NewDecoder(r.Body).Decode(&user)
	userSave, err := service.ActualizarUsuario(&user)
	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userSave)
}

func (controller *UsuarioController) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tiṕoId := chi.URLParam(r, "tipo")
	err := service.EliminarUsuario(id, tiṕoId)
	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(nil)
}

func (controller *UsuarioController) ValidatePass(w http.ResponseWriter, r *http.Request) {
	var user entidades.Usuario
	json.NewDecoder(r.Body).Decode(&user)
	role, err := service.ValidatePass(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(role)

}

func (controller *UsuarioController) ListarCorreo(w http.ResponseWriter, r *http.Request) {
	role := chi.URLParam(r, "id")
	email, err := service.ListarCorreo(role)
	if err != nil {
		http.Error(w, "Rol inválido", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(email)

}

func (controller *UsuarioController) ListarCorreoGestor(w http.ResponseWriter, r *http.Request) {
	role := chi.URLParam(r, "id")
	email, err := service.ListarCorreoGestor(role)
	if err != nil {
		http.Error(w, "Rol inválido", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(email)

}
