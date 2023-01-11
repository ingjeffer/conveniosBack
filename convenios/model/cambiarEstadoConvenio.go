package model

type CambiarEstadoConvenio struct {
	CambioEstado bool   `json:"cambioEstado"`
	Observacion  string `json:"observacion"`
}
