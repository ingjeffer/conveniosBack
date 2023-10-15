package model

import "encoding/json"

type Role string

const (
	Admin                           Role = "admin"
	Gestor                          Role = "gestor"
	Vicerectoria                    Role = "vicerectoria"
	Directo_Juridico                Role = "director juridico"
	Rectoria                        Role = "rectoria"
	Secretaria                      Role = "secretaria"
	Director_Relex                  Role = "director relex"
	Consejo_Academico               Role = "consejo academico"
	Consejo_Academico_Investigacion Role = "consejo academico investigacion"
	Director_Relex_Macro            Role = "director relex macro"
)

func (r Role) String() string {
	return string(r)
}

func (r *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
