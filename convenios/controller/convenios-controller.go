package controller

import (
	"convenios/model"
	"convenios/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type ConveniosController struct {
}

func (controller *ConveniosController) CrearConvenio(w http.ResponseWriter, r *http.Request) {
	var convenio model.Convenio
	json.NewDecoder(r.Body).Decode(&convenio)

	idHeader := r.Header.Get("x-id")
	convenio.IdGestorCreador = idHeader
	convenio.HistorialFirma = append(convenio.HistorialFirma, idHeader)
	convenioResp, err := service.GuardarConvenio(&convenio)

	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(convenioResp)
}

func (controller *ConveniosController) GetConvenios(w http.ResponseWriter, r *http.Request) {
	convenios, err := service.GetConvenios(r.Header.Get("x-role"), r.Header.Get("x-id"))

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

func (controller *ConveniosController) CambiarEstado(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	var cambio model.CambiarEstadoConvenio
	json.NewDecoder(r.Body).Decode(&cambio)

	if err := service.CambiarEstadoConvenio(id, cambio, r.Header.Get("x-role"), r.Header.Get("x-id")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
