package controller

import (
	"encoding/json"
	"net/http"
	"usuarios/service"
)

type RolesController struct {
}

func (controller *RolesController) ListarRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := service.ListarRoles()

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}
