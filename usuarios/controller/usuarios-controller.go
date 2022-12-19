package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"usuarios/entidades"
	"usuarios/model"
	"usuarios/service"
)

type UsuarioController struct {
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
	err := service.ValidatePass(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(nil)

}
