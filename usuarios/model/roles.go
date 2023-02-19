package model

type Roles struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre,omitempty" gorm:"column:nombre"`
}
