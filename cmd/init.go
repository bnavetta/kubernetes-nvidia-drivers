package cmd

import (
	"fmt"
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"github.com/spf13/cobra"
)

var InitCommand = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		vol, err := flexvol.Lookup()
		if err != nil {
			return err
		}

		err = vol.Create()
		if err != nil {
			return err
		}

		flexvol.Log(flexvol.Success(fmt.Sprintf("Initialized %v", vol)))

		return nil
	},
}
