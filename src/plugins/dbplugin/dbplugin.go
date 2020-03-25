package dbplugin

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"plugins_design_in_go/src/models"
)

type DbPlugin struct {
	pluginName string
	Repository
}

func NewDbPlugin() *DbPlugin {
	return &DbPlugin{}
}

func (p *DbPlugin) Name() string {
	return p.pluginName
}

func (p *DbPlugin) Initialize(ctx context.Context) error {
	p.pluginName = ctx.Value(models.DatabasePluginKey).(string)
	dbHost := ctx.Value(models.DB_HOST).(string)
	dbConnString := fmt.Sprintf("%s%s%s", "u:p@(",dbHost,":3306)/test?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(dbConnString)
	db, err := gorm.Open("mysql", dbConnString)
	if err != nil {
		log.Println(err)
		return err
	}
	p.db = migrateSchemas(db)
	return nil
}

func migrateSchemas(db *gorm.DB) *gorm.DB{
	return db.AutoMigrate(&models.Payment{}, &models.Product{})
}


func (p *DbPlugin) Stop() error {
	return p.db.Close()
}

