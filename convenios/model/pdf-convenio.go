package model

type ConvenioPDF struct {
	Convenio
	NumeroConvenio string
	FirmaInfo      []FirmaInfo
}

type FirmaInfo struct {
	Id        string
	Nombres   string
	Apellidos string
	TipoId    string
	Firma     string
	Rol       *Roles
}

type Roles struct {
	Nombre string
}
