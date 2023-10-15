package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convenio struct {
	ID                primitive.ObjectID `json:"id"`
	NombreInstitucion string             `json:"nombreInstitucion"`
	NombreConvenio    string             `json:"nombreConvenio"`
	ObjetoConvenio    string             `json:"objetoConvenio"`
	TipologiaConvenio string             `json:"tipologiaConvenio"`
	ModalidadConvenio string             `json:"modalidadConvenio"`
	Beneficiarios     string             `json:"beneficiarios"`
	Caracterizacion   string             `json:"caracterizacion"`
	InfoGestor        InfoGestor         `json:"infoGestor"`
	Estado            EstadoConvenio     `json:"estado"`
	Observaciones     string             `json:"observaciones,omitempty"`
	FirmaUrl          string             `json:"-"`
	IdGestorCreador   string             `json:"-"`
	HistorialFirma    []string           `json:"-"`
}

type InfoGestor struct {
	NombreResponsable string    `json:"nombreResponsable"`
	Fecha             time.Time `json:"fecha"`
	UnidadAcademica   string    `json:"unidadAcademica"`
	Cargo             string    `json:"cargo"`
	Email             string    `json:"email"`
	Telefono          string    `json:"telefono"`
}
