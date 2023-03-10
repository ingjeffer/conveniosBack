package controller

import (
	"convenios/model"
	"convenios/service"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type ConveniosController struct {
}

func (controller *ConveniosController) CrearConvenio(w http.ResponseWriter, r *http.Request) {
	var convenio model.Convenio
	json.NewDecoder(r.Body).Decode(&convenio)

	idHeader := r.Header.Get("x-id")
	convenio.IdGestorCreador = idHeader
	convenioResp, err := service.GuardarConvenio(&convenio)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(convenioResp)
}

func (controller *ConveniosController) GetConvenios(w http.ResponseWriter, r *http.Request) {
	convenios, err := service.GetConvenios()

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convenios)
}

func (controller *ConveniosController) GetConvenio(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	resp, err := service.GetConvenio(id)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (controller *ConveniosController) ActualizarConvenio(w http.ResponseWriter, r *http.Request) {
	var convenio model.Convenio
	json.NewDecoder(r.Body).Decode(&convenio)

	err := service.ActualizarConvenio(&convenio)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convenio)
}

func (controller *ConveniosController) GenerarPDFConvenio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	bytes, err := service.GenerarPDF(id)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=convenio.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

func (controller *ConveniosController) FirmarConvenio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("firma")

	if err != nil {
		http.Error(w, "Error campo firma", http.StatusBadRequest)
		return
	}

	if err = service.FirmarConvenio(id, file, handler); err != nil {
		http.Error(w, "Error al firmar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (controller *ConveniosController) CambiarEstado(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	var cambio model.CambiarEstadoConvenio
	json.NewDecoder(r.Body).Decode(&cambio)

	if err := service.CambiarEstadoConvenio(id, cambio, r.Header.Get("x-role")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
