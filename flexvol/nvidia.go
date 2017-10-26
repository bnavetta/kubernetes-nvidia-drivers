package flexvol

import (
	"fmt"
	"github.com/NVIDIA/nvidia-docker/src/nvidia"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func CreateVolume(mountPath string) (msg string, err error) {
	sentinelFile, err := os.Create("/home/ben/kube-nvidia-sentinel")
	if err != nil {
		return
	}
	defer sentinelFile.Close()

	base, err := ioutil.TempDir("", "kubernetes-nvidia-drivers")
	defer func() {
		if base != "" {
			os.RemoveAll(base)
		}
	}()

	if err != nil {
		return
	}

	fmt.Fprintf(sentinelFile, "Extracting volume into %s\n", base)

	vols, err := nvidia.LookupVolumes(base)
	if err != nil {
		return
	}

	driverVol := vols["nvidia_driver"]
	fmt.Fprintf(sentinelFile, "Found volume for version %s\n", driverVol.Version)
	err = driverVol.Create(nvidia.LinkOrCopyStrategy{})
	if err != nil {
		return
	}

	driverPath := path.Join(driverVol.Path, driverVol.Version)

	err = os.MkdirAll(filepath.Dir(mountPath), 0777)
	if err != nil {
		return
	}

	fmt.Fprintf(sentinelFile, "Renaming %s to %s\n", driverPath, mountPath)
	err = os.Rename(driverPath, mountPath)

	if err == nil {
		msg = fmt.Sprintf("Mounted volume v%s NVIDIA drivers at %s", driverVol.Version, mountPath)
	}

	return
}
