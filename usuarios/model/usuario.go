package model

type Usuario struct {
	Id        string `json:"id"`
	Nombres   string `json:"nombres,omitempty"`
	Apellidos string `json:"apellidos,omitempty"`
	TipoId    string `json:"tipoId,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Firma     string `json:"firma,omitempty"`
	Roles     *Roles `json:"rol,omitempty"`
}

type UsuarioCreate struct {
	RoleId int `json:"roleId,omitempty"`
	Usuario
}
