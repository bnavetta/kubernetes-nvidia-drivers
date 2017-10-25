package cmd

import (
	"errors"
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"github.com/spf13/cobra"
)

var UnmountCommand = &cobra.Command{
	Use: "unmount",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Usage: unmount <mount dir>")
		}

		mountDir := args[0]

		err := flexvol.Unbind(mountDir)
		if err != nil {
			return err
		}

		return nil
	},
}
