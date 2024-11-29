package shifu

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-github/github"
)

func GetInstallYaml(version string) (string, error) {
	yamlURL := "https://raw.githubusercontent.com/Edgenesis/shifu/refs/tags/" + version + "/pkg/k8s/crd/install/shifu_install.yml"
	resp, err := http.Get(yamlURL)
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
