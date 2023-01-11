package model

type EstadoConvenio string

const (
	Creado               EstadoConvenio = "CREADO"
	Firmado                             = "FIRMADO"
	Aprobado_Secretaria                 = "APROBADO_SECRETARIA"
	Rechazado_Secretaria                = "RECHAZADO_SECRETARIA"
)
