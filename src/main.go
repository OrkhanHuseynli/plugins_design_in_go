package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"plugins_design_in_go/src/app"
	"plugins_design_in_go/src/models"
	"plugins_design_in_go/src/plugins/controllerplugin"
	"plugins_design_in_go/src/plugins/dbplugin"

	database "plugins_design_in_go/src/repository"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"sample_graphql_in_go/src/controller"
	//"sample_graphql_in_go/src/graphqlapi"
	//database "sample_graphql_in_go/src/repository"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {

	//Query
	//query := `
	//	{
	//		payment(author: "Mitchel") {
	//				author
	//			}
	//	}
	//`

	//payments := dbManager.GetPaymentsByAuthor("Edo")

	//gqlprocessor := graphqlapi.NewGraphQLProcessor()
	//gqlprocessor.ExecuteQuery(query)
	//port := "5000"
	//service := controller.NewService(port)
	//service.Run()
	dbManager := database.NewDatabase()
	defer dbManager.Database.Close()

	app := app.New("plugins_app")
	dbPlugin := dbplugin.NewDbPlugin()
	servicepPlugin := controllerplugin.NewControllerPlugin(dbPlugin)
	fmt.Println("******* REGISTERING PLUGINS *******")
	app.Register(dbPlugin)
	app.Register(servicepPlugin)

	ctx, ctxCancel := context.WithCancel(context.Background())
	ctx = context.WithValue(context.Background(), models.ServiceNameKey, "ServiceNameKey")
	ctx = context.WithValue(context.Background(), models.DatabasePluginNameKey, "DatabasePluginNameKey")
	ctx = context.WithValue(ctx, models.DatabasePluginNameKey, "DB Plugin")
	ctx = context.WithValue(ctx, models.DB_HOST, "localhost")
	app.Start(ctx, ctxCancel)
}