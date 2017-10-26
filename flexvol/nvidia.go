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
	base, err := ioutil.TempDir("", "kubernetes-nvidia-drivers")
	defer func() {
		if base != "" {
			os.RemoveAll(base)
		}
	}()

	if err != nil {
		return
	}

	vols, err := nvidia.LookupVolumes(base)
	if err != nil {
		return
	}

	driverVol := vols["nvidia_driver"]
	err = driverVol.Create(nvidia.LinkOrCopyStrategy{})
	if err != nil {
		return
	}

	driverPath := path.Join(driverVol.Path, driverVol.Version)

	err = os.MkdirAll(filepath.Dir(mountPath), 0777)
	if err != nil {
		return
	}

	err = os.Rename(driverPath, mountPath)

	//contents, err := ioutil.ReadDir(driverPath)
	//if err != nil {
	//	return
	//}
	//
	//for _, entry := range contents {
	//	err = os.Rename(path.Join(driverPath, entry.Name()), path.Join(mountPath, entry.Name()))
	//	if err != nil {
	//		return
	//	}
	//}

	if err == nil {
		msg = fmt.Sprintf("Mounted volume v%s NVIDIA drivers at %s", driverVol.Version, mountPath)
	}

	return
}
