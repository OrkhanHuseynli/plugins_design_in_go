package controllerplugin

import (
	"context"
	"fmt"
	"net/http"
	"plugins_design_in_go/src/models"
	"plugins_design_in_go/src/plugins/dbplugin"
)

type ControllerPlugin struct {
	pluginName string
	dbPlugin dbplugin.IDbPlugin
	ctx context.Context
	server *http.Server
	//cancel context.CancelFunc
}

func NewControllerPlugin(dbPlugin dbplugin.IDbPlugin) *ControllerPlugin {
	return &ControllerPlugin{dbPlugin: dbPlugin}
}

func (p *ControllerPlugin) Name() string {
	return p.pluginName
}

func (p *ControllerPlugin) Initialize(ctx context.Context) error {
	//p.ctx, p.cancel  = context.WithCancel(ctx)
	p.pluginName = ctx.Value(models.ServicePluginNameKey).(string)
	fmt.Printf("Starting %s \n", p.pluginName)
	port := ctx.Value(models.ServicePortNumber).(string)
	handler := NewSimpleHandler(p.dbPlugin)
	//http.Handle("/payment", handler)
	//log.Printf("Server starting on port %v\n", port)
	//log.Fatal(http.ListenAndServe(":" + port, nil))

	router := http.NewServeMux()
	// Register your routes
	router.HandleFunc("/payment", handler.ServeHTTP)

	listenAddr := ":" + port

	p.server = &http.Server{
		Addr:         listenAddr,
		Handler:      router,
	}

	if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen on %s: %v\n", listenAddr, err)
		return err
	}
	return nil
}

func (p *ControllerPlugin) Stop() error {
	//p.cancel()
	return p.server.Shutdown(p.ctx)
}
