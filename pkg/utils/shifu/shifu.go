package shifu

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rhoninl/sft/pkg/k8s"
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
	return GetLatestShifuVersion()
}

func (shifu Shifu) SetVersion(version string) component {
	shifu.version = version

	return shifu
}

func (shifu Shifu) ResourceURL() string {
	if shifu.version == "" {
		shifu.version = GetLatestShifuVersion()
	}

	return fmt.Sprintf(shifu.basedUrl, shifu.version)
}

func (shifu Shifu) GetDeployYaml() (string, error) {
	return fetch(shifu.ResourceURL())
}

func fetch(url string) (string, error) {
	resp, err := http.Get(string(url))
	if err != nil {
		fmt.Println("Failed to install shifu component")
		return "", err
	}
	defer resp.Body.Close()

	yamlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to install shifu component")
		return "", err
	}

	return string(yamlContent), nil
}

func CheckShifuInstalled() error {
	if err := k8s.CheckCRDExists("edgedevices.shifu.edgenesis.io"); err != nil {
		return ErrorShifuUninstalled
	}

	return nil
}
