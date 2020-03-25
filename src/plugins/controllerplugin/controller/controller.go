package controller

import (
	"fmt"
	"github.com/plugins_design_in_go/src/plugins/dbplugin"
	"log"
	"net/http"
)

type Controller struct {
	port string
	dbPlugin *dbplugin.Plugin
}

func NewController(port string, dbPlugin *dbplugin.Plugin) *Controller {
	return &Controller{port: port, dbPlugin:dbPlugin}
}

func (c *Controller)Run()  {
	handler := NewSimpleHandler(c.dbPlugin)
	http.Handle("/payment", handler)
	log.Printf("Server starting on port %v\n", c.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", c.port), nil))
}