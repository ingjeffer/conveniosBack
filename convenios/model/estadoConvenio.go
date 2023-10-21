package model

type EstadoConvenio string

const (
	Creado                         EstadoConvenio = "CREADO_GESTOR"
	Firmado                        EstadoConvenio = "FIRMADO"
	Aprobado_Secretaria            EstadoConvenio = "APROBADO_SECRETARIA"
	Rechazado_Secretaria           EstadoConvenio = "RECHAZADO_SECRETARIA"
	Aprobado_Director_Relex        EstadoConvenio = "APROBADO_DIRECTOR_RELEX"
	Rechazado_Director_Relex       EstadoConvenio = "RECHAZADO_DIRECTOR_RELEX"
	Aprobado_Consejo_Academico_Inv EstadoConvenio = "APROBADO_CONSEJO_ACADEMICO"
	Aprobado_Consejo_Academico     EstadoConvenio = "APROBADO_CONSEJO_ACADEMICO"
	Rechazado_Consejo_Academico    EstadoConvenio = "RECHAZADO_CONSEJO_ACADEMICO"
	Rechazado_Vicerectoria         EstadoConvenio = "RECHAZADO_VICERECTORIA"
	Aprobado_Vicerectoria          EstadoConvenio = "APROBADO_VICERECTORIA"
	Aprobado_Director_Juridico     EstadoConvenio = "APROBADO_DIRECTOR_JURIDICO"
	Rechazado_Director_Juridico    EstadoConvenio = "RECHAZADO_DIRECTOR_JURIDICO"
	Aprobado_Director_Relex_Macro  EstadoConvenio = "APROBADO_DIRECTOR_RELEX"
	Aprobado_Rectoria              EstadoConvenio = "APROBADO_RECTORIA"
	Rechazado_Rectoria             EstadoConvenio = "RECHAZADO_RECTORIA"
	Convenio_Aprobado              EstadoConvenio = "APROBADO"
)
