package dbplugin

import (
	"context"
	"github.com/stretchr/testify/assert"
	"plugins_design_in_go/src/models"
	"testing"
)

func TestDbPlugin(t *testing.T) {
	dbplugin := NewDbPlugin()
	product := models.Product{Name:"Paste", Type:"Bathroom"}
	payment := models.Payment{"Unknown", "Sorbitol", product, "500" }
	dbplugin.IRepository = &MockRepository{nil, payment}
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	ctx = context.WithValue(ctx, models.DatabasePluginNameKey, "DB Plugin")
	ctx = context.WithValue(ctx, models.DB_HOST, "localhost")
	ctx = context.WithValue(ctx, models.DB_PORT, "3307")
	ctx = context.WithValue(ctx, models.DB_NAME, "test")
	ctx = context.WithValue(ctx, models.DB_USER, "user")
	ctx = context.WithValue(ctx, models.DB_PWD, "userPwd")
	err := dbplugin.Initialize(ctx)
	assert.Equal(t, "dial tcp [::1]:3307: connectex: No connection could be made because the target machine actively refused it.", err.Error())
}
