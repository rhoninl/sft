package shifu

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
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
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "Edgenesis", "shifu", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var version string
	for _, release := range releases {
		if strings.Contains(*release.Name, "rc") {
			continue
		}

		version = *release.Name
		break
	}
	return version
}

func GetAllAvailableVersions() []string {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "Edgenesis", "shifu", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var version []string

	for _, release := range releases {
		version = append(version, *release.Name)
	}

	return version
}
