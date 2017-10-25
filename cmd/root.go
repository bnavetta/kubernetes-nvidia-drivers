package cmd

import (
	"github.com/NVIDIA/nvidia-docker/src/nvml"
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"github.com/spf13/cobra"
	"os"
)

var RootCommand = &cobra.Command{
	Use:           "nvidiaDrivers",
	Short:         "Kubernetes FlexVolume driver for nvidia device driver libraries",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		os.Exit(1)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := nvml.Init(); err != nil {
			flexvol.Log(flexvol.Failure(err.Error()))
			os.Exit(1)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		nvml.Shutdown()
	},
}

func init() {
	RootCommand.AddCommand(InitCommand)
	RootCommand.AddCommand(MountCommand)
	RootCommand.AddCommand(UnmountCommand)
	RootCommand.AddCommand(GetVolumeNameCommand)
	RootCommand.AddCommand(Demo)
}
