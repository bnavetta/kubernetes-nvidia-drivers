package cmd

import (
	"fmt"
	"github.com/NVIDIA/nvidia-docker/src/nvidia"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

var Demo = &cobra.Command{
	Use: "demo",
	RunE: func(cmd *cobra.Command, args []string) error {
		prefix, err := filepath.Abs("kube-nvidia-drivers")
		if err != nil {
			return err
		}

		vols, err := nvidia.LookupVolumes(prefix)
		if err != nil {
			return err
		}

		driverVol := vols["nvidia_driver"]
		fmt.Printf("Volume %s at %s (version %s)\n", driverVol.Name, driverVol.Path, driverVol.Version)
		fmt.Printf("Mountpoint is %s, with options %s\n", driverVol.Mountpoint, driverVol.MountOptions)

		err = driverVol.Create(nvidia.LinkOrCopyStrategy{})
		if err != nil {
			return err
		}

		curPath := path.Join(driverVol.Path, driverVol.Version)
		volPath := path.Join(prefix, "current-driver")
		err = os.Rename(curPath, volPath)
		if err != nil {
			return err
		}

		fmt.Printf("Created volume contents at %s\n", volPath)

		return nil
	},
}
