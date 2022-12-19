package entidades

type Usuario struct {
	Id        string `gorm:"primaryKey; not null"`
	Nombres   string
	Apellidos string
	TipoId    string `gorm:"primaryKey; not null"`
	Email     string `gorm:"unique"`
	Password  string
	RoleId    int
	Roles     *Roles `gorm:"foreignKey:RoleId; omitempty"`
}
