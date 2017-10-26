package main

import (
	"fmt"
	"github.com/NVIDIA/nvidia-docker/src/nvidia"
	"github.com/NVIDIA/nvidia-docker/src/nvml"
	"os"
	"path"
)

const VolumeRoot = "/var/lib/kubernetes-nvidia-drivers"

func createVolume() error {
	err := nvml.Init()
	if err != nil {
		return fmt.Errorf("error initializing NVML: %s", err)
	}
	defer nvml.Shutdown()

	vols, err := nvidia.LookupVolumes(VolumeRoot)
	if err != nil {
		return fmt.Errorf("error looking up NVIDIA driver volumes: %s", err)
	}

	driverVol := vols["nvidia_driver"]
	fmt.Printf("Found volume for driver version %s in %s\n", driverVol.Version, driverVol.Path)

	err = driverVol.Create(nvidia.LinkOrCopyStrategy{})
	if err != nil {
		return fmt.Errorf("unable to extract driver files: %s", err)
	}

	driverPath := path.Join(driverVol.Path, driverVol.Version)
	currentPath := path.Join(VolumeRoot, "current")
	fmt.Printf("Linking %s into %s\n", driverPath, currentPath)
	err = os.Link(driverPath, currentPath)
	if err != nil {
		return fmt.Errorf("unable to link current driver version: %s", err)
	}

	return nil
}

func main() {
	if err := createVolume(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
