package main

import (
	"github.com/NVIDIA/nvidia-docker/src/nvml"
	"github.com/roguePanda/kubernetes-nvidia-drivers/cmd"
	"github.com/roguePanda/kubernetes-nvidia-drivers/flexvol"
	"os"
	"strings"
)

func main() {
	if err := nvml.Init(); err != nil {
		flexvol.Log(flexvol.Failure(err.Error()))
	}
	defer nvml.Shutdown()

	if err := cmd.RootCommand.Execute(); err != nil {
		if strings.HasPrefix(err.Error(), "unknown command") {
			flexvol.Log(flexvol.Reply{
				Status: flexvol.StatusNotSupported,
			})
		} else {
			flexvol.Log(flexvol.Failure(err.Error()))
			os.Exit(1)
		}
	}
}
