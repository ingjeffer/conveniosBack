package repository

import (
	"context"
	"fmt"
	"usuarios/configuration"
)

type ISessionRepository interface {
	GuardarSession()
}

var ctx = context.Background()

func GuardarSession(id string, token string, duration int64) {

	err := configuration.InstanceRedis.Set(ctx, "Id"+id, token, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sessión guardada correctamente -> ", "Id"+id)
	}
}
