package configuration

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"usuarios/entidades"
)

var Instance *gorm.DB

func IniciarDB() {

	var err error
	dns := "root:convenios@tcp(127.0.0.1:3306)/usuarios?charset=utf8mb4&parseTime=True&loc=Local"
	Instance, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		panic("Error al conectarse a la DB")
	}
	Instance.AutoMigrate(&entidades.Usuario{}, &entidades.Roles{})
	fmt.Println("Se ha cargado correctamente la config de BD")
}
