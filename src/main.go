package main

import (
	"context"
	"github.com/jinzhu/gorm"
	"plugins_design_in_go/src/app"
	"plugins_design_in_go/src/plugins/dbplugin"
	"plugins_design_in_go/src/plugins/controllerplugin"
	database "plugins_design_in_go/src/repository"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"sample_graphql_in_go/src/controller"
	//"sample_graphql_in_go/src/graphqlapi"
	"sample_graphql_in_go/src/models"
	//database "sample_graphql_in_go/src/repository"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {

	//Query
	query := `
		{
			payment(author: "Mitchel") {
					author
				}
		}
	`

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
	servicepPlugin := controllerplugin.NewServicePlugin(dbPlugin)
	app.Register(dbPlugin)
	app.Register(servicepPlugin)

	ctx, ctxCancel := context.WithCancel(context.Background())
	ctx = context.WithValue(context.Background(), models.ServiceNameKey, "ServiceNameKey")
	ctx = context.WithValue(context.Background(), models.DatabasePluginKey, "DatabasePluginKey")
	ctx = context.WithValue(ctx, models.ServiceNameKey, "ServiceNameKey")
	app.Start(ctx, ctxCancel)
	dbManager.AddPayment(&models.Payment{Author:"Edo", Product:"router", Sum:"400"})
}