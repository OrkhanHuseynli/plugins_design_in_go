package dbplugin

import (
	"context"
	"github.com/jinzhu/gorm"
	"log"
	"sample_graphql_in_go/src/models"
)

type Plugin struct {
	pluginName string
	Repository
}

func NewDbPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Name() string {
	return p.pluginName
}

func (p *Plugin) Initialize(ctx context.Context) error {
	p.pluginName = ctx.Value(models.DatabasePluginKey).(string)
	db, err := gorm.Open("mysql", "user:password@(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		return err
	}
	p.db = migrateSchemas(db)
	return nil
}

func migrateSchemas(db *gorm.DB) *gorm.DB{
	return db.AutoMigrate(&models.Payment{})
}


func (p *Plugin) Stop() error {
	return p.db.Close()
}

