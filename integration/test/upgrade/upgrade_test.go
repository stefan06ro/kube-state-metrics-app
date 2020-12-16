// +build k8srequired

package upgrade

import (
	"context"
	"testing"

	"github.com/giantswarm/apptest"

	"github.com/giantswarm/kube-state-metrics-app/integration/env"
	"github.com/giantswarm/kube-state-metrics-app/integration/key"
	"github.com/giantswarm/kube-state-metrics-app/integration/setup"
)

var (
	config setup.Config
)

func init() {
	var err error

	{
		config, err = setup.NewConfig()
		if err != nil {
			panic(err.Error())
		}
	}
}

func TestUpgrade(t *testing.T) {
	var err error
	ctx := context.Background()

	currentApp := apptest.App{
		CatalogName:   key.DefaultCatalogName(),
		Name:          key.CRName(),
		Namespace:     key.Namespace(),
		WaitForDeploy: true,
	}

	desiredApp := apptest.App{
		CatalogName:   key.DefaultTestCatalogName(),
		Name:          key.CRName(),
		Namespace:     key.Namespace(),
		SHA:           env.CircleSHA(),
		WaitForDeploy: true,
	}

	{
		err = config.AppTest.UpgradeApp(ctx, currentApp, desiredApp)
		if err != nil {
			t.Fatalf("expected nil got %#q", err)
		}
	}
}
