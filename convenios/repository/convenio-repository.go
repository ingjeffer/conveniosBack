package repository

import (
	"context"
	"convenios/configuration"
	"convenios/entidades"
	"convenios/model"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveConvenio(convenio *entidades.Convenio) (*entidades.Convenio, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := configuration.MongoC.Database("convenios-uis").Collection("convenios")
	convenio.Estado = model.Creado
	result, err := col.InsertOne(ctx, convenio)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	convenio.ID = result.InsertedID.(primitive.ObjectID)

	return convenio, nil
}

func GetConvenios(idGestor string, permisos []model.EstadoConvenio) ([]entidades.Convenio, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := configuration.MongoC.Database("convenios-uis").Collection("convenios")

	var filtro interface{}

	if idGestor != "" {
		filtro = bson.M{"idGestorCreador": idGestor}
	} else {
		filtro = bson.M{"estado": bson.M{"$in": permisos}}
	}

	result, err := col.Find(ctx, filtro)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var response []entidades.Convenio

	for result.Next(context.TODO()) {
		var elem entidades.Convenio
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		response = append(response, elem)
	}

	return response, nil
}

func GetConvenio(id string) (*entidades.Convenio, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	col := configuration.MongoC.Database("convenios-uis").Collection("convenios")

	var entity entidades.Convenio
	idSearch, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = col.FindOne(ctx, bson.D{{"_id", idSearch}}).Decode(&entity)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &entity, nil
}

func ActualizarConvenio(entidad *entidades.Convenio) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	col := configuration.MongoC.Database("convenios-uis").Collection("convenios")

	filter := bson.D{{"_id", entidad.ID}}

	if _, err := GetConvenio(entidad.ID.Hex()); err != nil {
		fmt.Println(err)
		return err
	}

	if _, err := col.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: entidad}}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
