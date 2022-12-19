package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
}

type InfoGestor struct {
	NombreResponsable string    `json:"nombreResponsable"`
	Fecha             time.Time `json:"fecha"`
	UnidadAcademica   string    `json:"unidadAcademica"`
	Cargo             string    `json:"cargo"`
	Email             string    `json:"email"`
	Telefono          string    `json:"telefono"`
}
