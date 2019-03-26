package startup

//go:generate go get github.com/golang/mock/mockgen
//go:generate mockgen -destination=../util/mocks/mock_$GOPACKAGE/startup.go -package=mock_$GOPACKAGE -source startup.go
//go:generate gofmt -s -l -w ../util/mocks/mock_$GOPACKAGE/startup.go
//go:generate go get golang.org/x/tools/cmd/goimports
//go:generate goimports -local=github.com/openshift/openshift-azure -e -w ../util/mocks/mock_$GOPACKAGE/startup.go

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	v4 "github.com/openshift/openshift-azure/pkg/startup/v4"
)

type Interface interface {
	WriteFiles(ctx context.Context) error
	Hash(role api.AgentPoolProfileRole) ([]byte, error)
}

func New(log *logrus.Entry, cs *api.OpenShiftManagedCluster) (Interface, error) {
	switch cs.Config.PluginVersion {
	case "v4.0":
		return v4.New(log, cs), nil
	}

	return nil, fmt.Errorf("version %q not found", cs.Config.PluginVersion)
}