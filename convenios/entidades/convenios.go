package entidades

import (
	"convenios/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convenio struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	NombreInstitucion string               `bson:"nombreInstitucion,omitempty" json:"nombreInstitucion"`
	NombreConvenio    string               `bson:"nombreConvenio,omitempty" json:"nombreConvenio"`
	ObjetoConvenio    string               `bson:"objetoConvenio,omitempty" json:"objetoConvenio"`
	TipologiaConvenio string               `bson:"tipologiaConvenio,omitempty" json:"tipologiaConvenio"`
	ModalidadConvenio string               `bson:"modalidadConvenio,omitempty" json:"modalidadConvenio"`
	Beneficiarios     string               `bson:"beneficiarios,omitempty" json:"beneficiarios"`
	Caracterizacion   string               `bson:"caracterizacion,omitempty" json:"caracterizacion"`
	InfoGestor        InfoGestor           `bson:"infoGestor,omitempty" json:"infoGestor"`
	Estado            model.EstadoConvenio `bson:"estado,omitempty" json:"estado"`
	FirmaUrl          string               `bson:"firmaUrl,omitempty" json:"firmaUrl"`
	Observaciones     string               `bson:"observaciones,omitempty" json:"observaciones"`
	IdGestorCreador   string               `bson:"idGestorCreador,omitempty" json:"idGestorCreador"`
	HistorialFirma    []string             `bson:"historialFirma,omitempty" json:"historialFirma"`
}

type InfoGestor struct {
	NombreResponsable string    `bson:"nombreResponsable,omitempty" json:"nombreResponsable"`
	Fecha             time.Time `bson:"fecha,omitempty" json:"fecha"`
	UnidadAcademica   string    `bson:"unidadAcademica,omitempty" json:"unidadAcademica"`
	Cargo             string    `bson:"cargo,omitempty" json:"cargo"`
	Email             string    `bson:"email,omitempty" json:"email"`
	Telefono          string    `bson:"telefono,omitempty" json:"telefono"`
}
