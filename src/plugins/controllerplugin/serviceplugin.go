package controllerplugin

import (
	"context"
	"github.com/plugins_design_in_go/src/models"
	"github.com/plugins_design_in_go/src/plugins/dbplugin"
)

type Plugin struct {
	pluginName string
	dbPlugin *dbplugin.Plugin
}

func NewServicePlugin(dbPlugin *dbplugin.Plugin) *Plugin {
	return &Plugin{dbPlugin:dbPlugin}
}

func (p *Plugin) Name() string {
	return p.pluginName
}

func (p *Plugin) Initialize(ctx context.Context) error {
	p.pluginName = ctx.Value(models.ServicePluginKey).(string)
	return nil
}

func (p *Plugin) Stop() error {
	return nil
}
