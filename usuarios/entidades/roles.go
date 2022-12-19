package entidades

type Roles struct {
	Id     int `gorm:"primaryKey; not null; omitempty"`
	Nombre string
}
