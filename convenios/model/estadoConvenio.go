package model

type EstadoConvenio string

const (
	Creado                         EstadoConvenio = "CREADO"
	Firmado                                       = "FIRMADO"
	Aprobado_Secretaria                           = "APROBADO_SECRETARIA"
	Rechazado_Secretaria                          = "RECHAZADO_SECRETARIA"
	Aprobado_Director_Relex                       = "APROBADO_DIRECTOR_RELEX"
	Rechazado_Director_Relex                      = "RECHAZADO_DIRECTOR_RELEX"
	Aprobado_Consejo_Academico_Inv                = "APROBADO_CONSEJO_ACADEMICO_INV"
	Aprobado_Consejo_Academico                    = "APROBADO_CONSEJO_ACADEMICO"
	Rechazado_Consejo_Academico                   = "RECHAZADO_CONSEJO_ACADEMICO"
	Aprobado_Vicerectoria                         = "APROBADO_VICERECTORIA"
	Aprobado_Director_Juridico                    = "APROBADO_DIRECTOR_JURIDICO"
	Rechazado                                     = "CONVENIO_RECHAZADO"
)
