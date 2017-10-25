package cmd

import (
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"github.com/spf13/cobra"
)

var GetVolumeNameCommand = &cobra.Command{
	Use: "getvolumename",
	RunE: func(cmd *cobra.Command, args []string) error {
		vol, err := flexvol.Lookup()
		if err != nil {
			return err
		}

		flexvol.Log(flexvol.Reply{
			Status:     flexvol.StatusSuccess,
			VolumeName: vol.Name(),
		})

		return nil
	},
}
