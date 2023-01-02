package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"usuarios/entidades"
)

func GenerarJWT(usuario *entidades.Usuario) (string, int64, error) {
	clave := []byte("convenios-uis")

	now := time.Now().Add(time.Minute * 10).Unix()

	payload := jwt.MapClaims{
		"role":     usuario.Roles,
		"id":       usuario.Id,
		"email":    usuario.Email,
		"nombre":   usuario.Nombres,
		"apellido": usuario.Apellidos,
		"exp":      now,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clave)
	if err != nil {
		fmt.Println(err.Error())
		return tokenStr, 0, err
	}
	return tokenStr, now, nil
}
