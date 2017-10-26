package main

import (
	"fmt"
	"github.com/NVIDIA/nvidia-docker/src/nvidia"
	"github.com/NVIDIA/nvidia-docker/src/nvml"
	"io/ioutil"
	"os"
	"path"
)

const VolumeRoot = "/var/lib/kubernetes-nvidia-drivers"

// Been getting weird hardlink issues
type CopyStrategy struct{}

func (s CopyStrategy) Clone(src, dst string) error {
	return nvidia.Copy(src, dst)
}

func createVolume() error {
	err := nvml.Init()
	if err != nil {
		return fmt.Errorf("error initializing NVML: %s", err)
	}
	defer nvml.Shutdown()

	tmp, err := ioutil.TempDir("", "kubernetes-nvidia-drivers")
	if err != nil {
		return fmt.Errorf("unable to create staging directory: %s", err)
	}
	defer os.RemoveAll(tmp)

	vols, err := nvidia.LookupVolumes(tmp)
	if err != nil {
		return fmt.Errorf("error looking up NVIDIA driver volumes: %s", err)
	}

	driverVol := vols["nvidia_driver"]
	fmt.Printf("Found volume for driver version %s in %s\n", driverVol.Version, driverVol.Path)

	err = driverVol.Create(CopyStrategy{})
	if err != nil {
		return fmt.Errorf("unable to extract driver files: %s", err)
	}

	driverPath := path.Join(driverVol.Path, driverVol.Version)
	finalPath := path.Join(VolumeRoot, "drivers")
	fmt.Printf("Renaming %s to %s\n", driverPath, finalPath)
	err = os.Rename(driverPath, finalPath)
	if err != nil {
		return fmt.Errorf("unable to set current driver version: %s", err)
	}

	return nil
}

func main() {
	if err := createVolume(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
