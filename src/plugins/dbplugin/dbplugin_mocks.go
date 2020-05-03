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
	dbConnString := fmt.Sprintf("%s%s%s%s%s", "u:u@(",dbHost,":", dbPort,")/test?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf(dbConnString)
	return nil
}

func (m *MockDbPlugin) Stop() error {
	return nil
}
