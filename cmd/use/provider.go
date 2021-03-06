package use

import (
	"strings"

	"github.com/devspace-cloud/devspace/pkg/devspace/cloud/config"
	"github.com/pkg/errors"

	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/devspace-cloud/devspace/pkg/util/survey"
	"github.com/spf13/cobra"
)

type providerCmd struct{}

func newProviderCmd() *cobra.Command {
	cmd := &providerCmd{}

	return &cobra.Command{
		Use:   "provider",
		Short: "Change the default provider",
		Long: `
#######################################################
############### devspace use provider #################
#######################################################
Use a specific cloud provider as default for future
commands

Example:
devspace use provider my.domain.com
#######################################################
	`,
		Args: cobra.MaximumNArgs(1),
		RunE: cmd.RunUseProvider,
	}
}

// RunUseProvider executes the "devspace use provider" command logic
func (*providerCmd) RunUseProvider(cobraCmd *cobra.Command, args []string) error {
	// Get provider configuration
	providerConfig, err := config.ParseProviderConfig()
	if err != nil {
		return errors.Errorf("Error loading provider config: %v", err)
	}

	providerName := ""
	if len(args) > 0 {
		providerName = args[0]
	} else {
		providerNames := make([]string, 0, len(providerConfig.Providers))
		for _, provider := range providerConfig.Providers {
			providerNames = append(providerNames, strings.TrimSpace(provider.Name))
		}

		providerName, err = survey.Question(&survey.QuestionOptions{
			Question: "Please select a default provider",
			Options:  providerNames,
		}, log.GetInstance())
		if err != nil {
			return err
		}
	}

	provider := config.GetProvider(providerConfig, providerName)
	if provider == nil {
		return errors.Errorf("Error provider %s does not exist! Did you run `devspace add provider %s` first?", providerName, providerName)
	}

	providerConfig.Default = provider.Name
	err = config.SaveProviderConfig(providerConfig)
	if err != nil {
		return errors.Errorf("Couldn't save provider config: %v", err)
	}

	log.Donef("Successfully changed default cloud provider to %s", providerName)
	return nil
}
