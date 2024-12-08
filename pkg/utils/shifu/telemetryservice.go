package shifu

import (
	"fmt"

	"github.com/rhoninl/sft/pkg/utils/cache"
	"github.com/rhoninl/sft/pkg/utils/logger"
)

const (
	telemetryServiceInstallYamlBaseURL = "https://raw.githubusercontent.com/Edgenesis/shifu/refs/tags/%s/pkg/telemetryservice/install/telemetryservice_install.yaml"
)

type TelemetryService struct {
	version  string
	basedUrl string
}

func newTelemetryService() TelemetryService {
	return TelemetryService{
		basedUrl: telemetryServiceInstallYamlBaseURL,
	}
}

func (telemetryService TelemetryService) Version() string {
	logger.Debugf(logger.MoreVerbose, "using telemetryservice version: %s", telemetryService.version)
	return GetLatestShifuVersion()
}

func (telemetryService TelemetryService) SetVersion(version string) component {
	telemetryService.version = version

	return telemetryService
}

func (telemetryService *TelemetryService) ResourceURL() string {
	if telemetryService.version == "" {
		telemetryService.version = GetLatestShifuVersion()
	}

	return fmt.Sprintf(telemetryService.basedUrl, telemetryService.version)
}

func (telemetryService TelemetryService) GetDeployYaml() (string, error) {
	cacherName := fmt.Sprintf("%s.%s", "ts", telemetryService.Version())
	data, err := cache.GetOrDoAndCache(cacherName, func() ([]byte, error) {
		return fetch(telemetryService.ResourceURL())
	})

	logger.Debugf(logger.MoreVerbose, "fetched telemetryservice yaml: %s", string(data))

	return string(data), err
}
