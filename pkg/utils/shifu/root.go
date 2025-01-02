package shifu

import (
	"context"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/rhoninl/sft/pkg/utils/cache"
	"github.com/rhoninl/sft/pkg/utils/logger"
)

type ResourceType uint8

const (
	TypeShifu ResourceType = iota + 1
	TypeTelemetryService
)

type component interface {
	Version() string
	SetVersion(version string) component
	GetDeployYaml() (string, error)
}

func Resource(resourceType ResourceType) component {
	switch resourceType {
	case TypeShifu:
		return newShifu()
	case TypeTelemetryService:
		return newTelemetryService()
	default:
		return nil
	}
}

func GetLatestShifuVersion() string {
	releases := GetAllAvailableVersions()

	logger.Debugf(logger.MoreVerbose, "got shifu releases: %v", releases)

	var version string
	for _, release := range releases {
		if strings.Contains(release, "rc") {
			continue
		}

		version = release
		break
	}

	logger.Debugf(logger.MoreVerbose, "latest shifu version: %s", version)
	return version
}

func GetAllAvailableVersions() []string {
	versions, err := cache.DoWithExpire("shifu_releases", time.Hour, func() ([]byte, error) {
		client := github.NewClient(nil)
		releases, _, err := client.Repositories.ListReleases(context.Background(), "Edgenesis", "shifu", nil)
		if err != nil {
			logger.Debugf(logger.Verbose, "Failed to get shifu releases: %v", err)
			return nil, err
		}

		var version []string

		for _, release := range releases {
			version = append(version, *release.Name)
		}

		return []byte(strings.Join(version, ",")), nil
	})

	if err != nil {
		logger.Debugf(logger.Verbose, "Failed to get shifu releases: %v", err)
		return nil
	}

	return strings.Split(string(versions), ",")
}
