package cmd

import (
	"encoding/json"
	"errors"
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var MountCommand = &cobra.Command{
	Use: "mount",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Usage: mount <mount dir> <json params>")
		}

		mountDir := args[0]

		f, _ := os.Create("/tmp/mnt-" + strings.Replace(mountDir, "/", "_", -1))
		b, _ := json.Marshal(args)
		f.Write(b)
		f.Close()

		err := os.MkdirAll(mountDir, 0755)
		if err != nil {
			return err
		}

		vol, err := flexvol.Lookup()
		if err != nil {
			return err
		}

		err = vol.Bind(mountDir)
		if err != nil {
			return err
		}

		return err
	},
}
