package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func DropProfile(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("must specifiy a profile name")
	}
	profile := profileFile(args[0])
	if _, err := os.Stat(profile); err != nil {
		return fmt.Errorf("profile '%s' not exists", args[0])
	}
	if err := os.Remove(profile); err != nil {
		log.Error().Msgf("Failed to remove profile '%s'", args[0])
		return err
	} else {
		log.Info().Msgf("Profile '%s' removed", args[0])
	}
	return nil
}

func DropProfileHandle(cmd *cobra.Command) {
	cmd.ValidArgsFunction = profileValidator
}
