package sync

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/openshift/openshift-azure/pkg/entrypoint/config"
)

type cmdConfig struct {
	config.Common
	dryRun          bool
	once            bool
	interval        time.Duration
	httpPort        int
	metricsEndpoint string
}

// NewCommand returns the cobra command for "sync".
func NewCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:  "sync",
		Long: "Start sync application",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := configFromCmd(cmd)
			if err != nil {
				return err
			}
			return start(cfg)
		},
	}
	cc.Flags().Bool("dry-run", false, "Print resources to be synced instead of mutating cluster state.")
	cc.Flags().Bool("run-once", false, "If true, run only once then quit.")
	cc.Flags().Duration("interval", 3*time.Minute, "How often the sync process going to be rerun.")
	cc.Flags().Int("http-port", 8080, "The http server port")
	cc.Flags().String("metrics-endpoint", "/metrics", "The endpoint for serving sync metrics")

	return cc
}

func configFromCmd(cmd *cobra.Command) (*cmdConfig, error) {
	c := &cmdConfig{}
	var err error
	c.Common, err = config.CommonConfigFromCmd(cmd)
	if err != nil {
		return nil, err
	}
	c.dryRun, err = cmd.Flags().GetBool("dry-run")
	if err != nil {
		return nil, err
	}
	c.once, err = cmd.Flags().GetBool("run-once")
	if err != nil {
		return nil, err
	}
	c.interval, err = cmd.Flags().GetDuration("interval")
	if err != nil {
		return nil, err
	}
	c.httpPort, err = cmd.Flags().GetInt("http-port")
	if err != nil {
		return nil, err
	}
	c.metricsEndpoint, err = cmd.Flags().GetString("metrics-endpoint")
	if err != nil {
		return nil, err
	}

	return c, nil
}
