package dbplugin

import (
	"context"
	"plugins_design_in_go/src/models"
	"testing"
)

//NOTE : these tests can run only with running DB
// you can run DB in the docker container `docker-compose up db`

func TestRepository_AddPayment(t *testing.T) {
	dbplugin := NewDbPlugin()
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	ctx = context.WithValue(ctx, models.DatabasePluginKey, "DB Plugin")
	ctx = context.WithValue(ctx, models.DB_HOST, "localhost")
	ctx = context.WithValue(ctx, models.DB_PORT, "3307")
	err := dbplugin.Initialize(ctx)
	if err != nil {
		panic("The database failed to init")
	}
	product1 := models.Product{Name: "mockname", Type:"food"}
	product2 := models.Product{Name: "vikname", Type: "food"}
	payment1 := models.Payment{Author:"Moris", Product:product1,Sum: "399"}
	payment2 := models.Payment{Author:"Dotis", Product:product2,Sum: "777"}
	dbplugin.AddPayment(&payment1)
	dbplugin.AddPayment(&payment2)
}

func TestRepository_DeletePaymentsByAuthor(t *testing.T) {
	author := "Moris"
	dbplugin := NewDbPlugin()
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	ctx = context.WithValue(ctx, models.DatabasePluginKey, "DB Plugin")
	ctx = context.WithValue(ctx, models.DB_HOST, "localhost")
	ctx = context.WithValue(ctx, models.DB_PORT, "3307")
	err := dbplugin.Initialize(ctx)
	if err != nil {
		panic("The database failed to init")
	}
	dbplugin.DeletePaymentsByAuthor(author)
}