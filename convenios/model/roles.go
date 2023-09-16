package model

import "encoding/json"

type Role string

const (
	Admin                 Role = "admin"
	Gestor                Role = "gestor"
	Vicerectoria          Role = "vicerectoria"
	Directo_Juridico      Role = "director juridico"
	Rectoria              Role = "rectoria"
	Secretaria            Role = "secretaria"
	Director_Relex        Role = "director relex"
	Consejo_Academico     Role = "consejo academico"
	Consejo_Academico_Inv Role = "consejo academico investigacion"
)

func (r Role) String() string {
	return string(r)
}

func (r *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
