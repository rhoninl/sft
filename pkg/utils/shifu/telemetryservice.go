package shifu

import "fmt"

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
	return GetLatestShifuVersion()
}

func (telemetryService TelemetryService) SetVersion(version string) component {
	telemetryService.version = version

	return telemetryService
}

func (telemetryService TelemetryService) ResourceURL() string {
	if telemetryService.version == "" {
		telemetryService.version = GetLatestShifuVersion()
	}

	return fmt.Sprintf(telemetryService.basedUrl, telemetryService.version)
}

func (telemetryService TelemetryService) GetDeployYaml() (string, error) {
	return fetch(telemetryService.ResourceURL())
}
