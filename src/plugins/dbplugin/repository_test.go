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
	product1 := models.Product{"mockname", "food"}
	product2 := models.Product{"vikname", "food"}
	payment1 := models.Payment{"Moris", product1, "399"}
	payment2 := models.Payment{"Moris", product2, "444"}
	dbplugin.AddPayment(&payment1)
	dbplugin.AddPayment(&payment2)
}
