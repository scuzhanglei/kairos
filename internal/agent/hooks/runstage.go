package hook

import (
	"github.com/kairos-io/kairos-sdk/utils"
	config "github.com/kairos-io/kairos/v2/pkg/config"

	events "github.com/kairos-io/kairos-sdk/bus"
)

type RunStage struct{}

func (r RunStage) Run(_ config.Config) error {
	utils.SH("elemental run-stage kairos-install.after")             //nolint:errcheck
	events.RunHookScript("/usr/bin/kairos-agent.install.after.hook") //nolint:errcheck
	return nil
}
