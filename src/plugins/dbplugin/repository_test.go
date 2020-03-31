package dbplugin

import (
	"context"
	"github.com/stretchr/testify/assert"
	"plugins_design_in_go/src/models"
	"testing"
)

//NOTE : these tests can run only with running DB
// you can run DB in the docker container `docker-compose up db`
const (
	author1 = "Moris"
	author2 = "Dotis"
)

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
	payment1 := models.Payment{Author:author1, Product:product1,Sum: "399"}
	payment2 := models.Payment{Author:author2, Product:product2,Sum: "777"}
	dbplugin.AddPayment(&payment1)
	dbplugin.AddPayment(&payment2)
	retrievedPayment1 := dbplugin.GetPaymentsByAuthor(author1)
	retrievedPayment2 := dbplugin.GetPaymentsByAuthor(author2)
	retrievedPayments := append(retrievedPayment1,retrievedPayment2... )
	for _, p := range retrievedPayments {
		if p.Author == author1 {
			assert.Equal(t, payment1.Sum, p.Sum, "The two sums should be the same.")
		} else if p.Author == author2 {
			assert.Equal(t, payment2.Sum, p.Sum, "The two sums should be the same.")
		}
	}

	retrievedPayments = dbplugin.GetPaymentsByAuthorEagerly(author1)
	for _, p := range retrievedPayments {
		if p.Author == author1 {
			assert.Equal(t, payment1.Product.Name, p.Product.Name, "The two product names should be the same.")
		} else if p.Author == author2 {
			assert.Equal(t, payment2.Product.Name, p.Product.Name, "The two product names should be the same.")
		}
	}
	dbplugin.DeletePaymentsByAuthor(author1)
	dbplugin.DeletePaymentsByAuthor(author2)
	retrievedPayments = dbplugin.GetPaymentsByAuthor(author1)
	assert.Equal(t, 0, len(retrievedPayments), "The two lens should be the same.")
	retrievedPayments = dbplugin.GetPaymentsByAuthor(author2)
	assert.Equal(t, 0, len(retrievedPayments), "The two lens should be the same.")
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