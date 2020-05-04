package controllerplugin

import (
	"context"
	"github.com/stretchr/testify/assert"
	"plugins_design_in_go/src/models"
	"plugins_design_in_go/src/plugins/dbplugin"
	"testing"
	"time"
)

type MockServer struct {
}

func (m MockServer) ListenAndServe() error {
	return nil
}
func (m MockServer) Shutdown(ctx context.Context) error {
	return nil
}

func TestControllerPlugin(t *testing.T) {
	product := models.Product{Name: "Paste", Type: "Bathroom"}
	payment := models.Payment{"Unknown", "Sorbitol", product, "500"}
	repo := dbplugin.MockRepository{MockObject: payment}
	dbpl := dbplugin.MockDbPlugin{IRepository: &repo}
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	ctx = context.WithValue(ctx, models.ServicePluginNameKey, "Service Plugin")
	ctx = context.WithValue(ctx, models.DatabasePluginNameKey, "DB Plugin")
	ctx = context.WithValue(ctx, models.DB_HOST, "db")
	ctx = context.WithValue(ctx, models.DB_PORT, "3307")
	ctx = context.WithValue(ctx, models.DB_NAME, "test")
	ctx = context.WithValue(ctx, models.DB_USER, "user")
	ctx = context.WithValue(ctx, models.DB_PWD, "userPwd")
	ctx = context.WithValue(ctx, models.ServicePortNumber, "5000")

	ctrpl := NewControllerPlugin(dbpl)
	ctrpl.server = MockServer{}
	err := ctrpl.Initialize(ctx)
	assert.NoError(t, err)
	time.Sleep(3 * time.Second)
	err = ctrpl.Stop()
	assert.NoError(t, err)
}
