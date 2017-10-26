package main

import (
	"github.com/NVIDIA/nvidia-docker/src/nvml"
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var InitCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		flexvol.Log(flexvol.Reply{
			Status:       flexvol.StatusSuccess,
			Capabilities: &flexvol.Capabilities{Attach: false},
		})
	},
}

var MountDeviceCmd = &cobra.Command{
	Use: "mountdevice",
	Run: func(cmd *cobra.Command, args []string) {
		mountDir := args[0]
		msg, err := flexvol.CreateVolume(mountDir)
		if err != nil {
			flexvol.Log(flexvol.Failure(err.Error()))
		} else {
			flexvol.Log(flexvol.Success(msg))
		}
	},
}

var UnmountDeviceCmd = &cobra.Command{
	Use: "unmountdevice",
	Run: func(cmd *cobra.Command, args []string) {
		mountDir := args[0]
		children, err := ioutil.ReadDir(mountDir)
		if err != nil {
			flexvol.Log(flexvol.Failure(err.Error()))
			return
		}

		for _, child := range children {
			err = os.RemoveAll(path.Join(mountDir, child.Name()))
			if err != nil {
				flexvol.Log(flexvol.Failure(err.Error()))
				return
			}
		}

		flexvol.Log(flexvol.Success("Successfully removed directory contents"))
	},
}

var RootCmd = &cobra.Command{
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

func main() {
	RootCmd.AddCommand(InitCmd, MountDeviceCmd, UnmountDeviceCmd)
	err := RootCmd.Execute()
	if err != nil {
		if strings.HasPrefix(err.Error(), "unknown command") {
			flexvol.Log(flexvol.Reply{
				Status: flexvol.StatusNotSupported,
			})
		} else {
			flexvol.Log(flexvol.Failure(err.Error()))
		}
	}
}
