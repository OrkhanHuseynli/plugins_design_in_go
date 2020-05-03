package controllerplugin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"plugins_design_in_go/src/models"
	"plugins_design_in_go/src/plugins/dbplugin"
)

type ControllerPlugin struct {
	pluginName string
	dbPlugin dbplugin.IDbPlugin
	ctx context.Context
	cancel context.CancelFunc
}

func NewControllerPlugin(dbPlugin dbplugin.IDbPlugin) *ControllerPlugin {
	return &ControllerPlugin{dbPlugin: dbPlugin}
}

func (p *ControllerPlugin) Name() string {
	return p.pluginName
}

func (p *ControllerPlugin) Initialize(ctx context.Context) error {
	p.ctx, p.cancel  = context.WithCancel(ctx)
	p.pluginName = ctx.Value(models.ServicePluginNameKey).(string)
	port := ctx.Value(models.ServicePortNumber).(string)
	handler := NewSimpleHandler(p.dbPlugin)
	http.Handle("/payment", handler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	return nil
}

func (p *ControllerPlugin) Stop() error {
	p.cancel()
	return nil
}
