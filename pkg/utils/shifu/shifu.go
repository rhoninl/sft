package shifu

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/cache"
	"github.com/rhoninl/sft/pkg/utils/logger"
)

const (
	shifuInstallYamlBaseURL = "https://raw.githubusercontent.com/Edgenesis/shifu/refs/tags/%s/pkg/k8s/crd/install/shifu_install.yml"
)

var (
	ErrorShifuUninstalled = errors.New("Shifu is not installed in the cluster")
)

type Shifu struct {
	basedUrl string
	version  string
}

func newShifu() Shifu {
	return Shifu{
		basedUrl: shifuInstallYamlBaseURL,
	}
}

func ShifuVersion(version string) Shifu {
	return Shifu{}
}

func (shifu Shifu) Version() string {
	logger.Debugf(logger.Verbose, "using shifu version: %s", shifu.version)
	return shifu.version
}

func (shifu Shifu) SetVersion(version string) component {
	if len(version) > 0 {
		shifu.version = version
	}

	return shifu
}

func (shifu *Shifu) ResourceURL() string {
	if shifu.version == "" || shifu.version == "latest" {
		shifu.version = GetLatestShifuVersion()
	}

	return fmt.Sprintf(shifu.basedUrl, shifu.version)
}

func (shifu Shifu) GetDeployYaml() (string, error) {
	url := shifu.ResourceURL()
	cacherName := fmt.Sprintf("%s.%s", "shifu", shifu.Version())
	data, err := cache.GetOrDoAndCache(cacherName, func() ([]byte, error) {
		return fetch(url)
	})
	logger.Debugf(logger.MoreVerbose, "fetched shifu yaml from %s: %s", url, string(data))

	return string(data), err
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(string(url))
	if err != nil {
		logger.Debugf(logger.Verbose, "Failed to install shifu component: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	yamlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Debugf(logger.Verbose, "Failed to install shifu component: %v", err)
		return nil, err
	}

	logger.Debugf(logger.MoreVerbose, "fetched shifu yaml: %s", string(yamlContent))

	return yamlContent, nil
}

func CheckShifuInstalled() error {
	if err := k8s.CheckCRDExists("edgedevices.shifu.edgenesis.io", "v1alpha1"); err != nil {
		return ErrorShifuUninstalled
	}

	logger.Debugf(logger.MoreVerbose, "shifu is installed")
	return nil
}

// func Delete(deviceName string) error {
// 	device, err := k8s.GetAllByDeviceName(deviceName)
// 	if err != nil {
// 		return fmt.Errorf("failed to get all by device name")
// 	}

// }
