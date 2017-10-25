package flexvol

import (
	"errors"
	"fmt"
	"github.com/NVIDIA/nvidia-docker/src/nvidia"
	"golang.org/x/sys/unix"
	"os"
	"path"
)

type DriverVolume struct {
	vol *nvidia.Volume
}

const NvidiaVolumePrefix = "/var/lib/kubernetes-nvidia-drivers"

func Lookup() (*DriverVolume, error) {
	vols, err := nvidia.LookupVolumes(NvidiaVolumePrefix)
	if err != nil {
		return nil, err
	}

	vol, exists := vols["nvidia_driver"]
	if !exists {
		return nil, errors.New("nvidia_driver volume not found")
	}

	return &DriverVolume{vol}, nil
}

func (d *DriverVolume) Create() error {
	err := d.vol.Create(nvidia.LinkStrategy{})
	if err != nil {
		//os.RemoveAll(d.vol.Path)
		return err
	}

	return nil
}

func (d *DriverVolume) Remove() error {
	return os.RemoveAll(d.vol.Path)
}

func (d *DriverVolume) Bind(target string) error {
	return unix.Mount(path.Join(d.vol.Path, d.vol.Version), target, "", unix.MS_BIND, "")
}

func Unbind(target string) error {
	return unix.Unmount(target, 0)
}

func (d *DriverVolume) Name() string {
	return fmt.Sprintf("%v_%v", d.vol.Name, d.vol.Version)
}

func (d *DriverVolume) String() string {
	return fmt.Sprintf("NVIDIA driver volume version %v at %v", d.vol.Version, d.vol.Path)
}
