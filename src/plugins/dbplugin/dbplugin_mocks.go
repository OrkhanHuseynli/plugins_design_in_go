package dbplugin

import (
	"context"
	"fmt"
	"plugins_design_in_go/src/models"
)

type MockDbPlugin struct {
	pluginName string
	IRepository
}

func (m *MockDbPlugin) Name() string {
	return "mockDbPlugin"
}

func (m *MockDbPlugin) Initialize(ctx context.Context) error {
	m.pluginName = ctx.Value(models.DatabasePluginNameKey).(string)
	dbHost := ctx.Value(models.DB_HOST).(string)
	dbPort := ctx.Value(models.DB_PORT).(string)
	dbName := ctx.Value(models.DB_NAME).(string)
	dbUser := ctx.Value(models.DB_USER).(string)
	dbPwd := ctx.Value(models.DB_PWD).(string)

	dbConnString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", dbUser,":",dbPwd,"@(",dbHost,":", dbPort,")/",dbName,"?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf(dbConnString)
	return nil
}

func (m *MockDbPlugin) Stop() error {
	return nil
}
