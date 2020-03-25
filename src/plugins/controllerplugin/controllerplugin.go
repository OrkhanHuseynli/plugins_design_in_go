package controllerplugin

import (
	"context"
	"plugins_design_in_go/src/models"
	"plugins_design_in_go/src/plugins/dbplugin"
)

type ControllerPlugin struct {
	pluginName string
	dbPlugin *dbplugin.DbPlugin
}

func NewControllerPlugin(dbPlugin *dbplugin.DbPlugin) *ControllerPlugin {
	return &ControllerPlugin{dbPlugin: dbPlugin}
}

func (p *ControllerPlugin) Name() string {
	return p.pluginName
}

func (p *ControllerPlugin) Initialize(ctx context.Context) error {
	p.pluginName = ctx.Value(models.ServicePluginKey).(string)
	return nil
}

func (p *ControllerPlugin) Stop() error {
	return nil
}
