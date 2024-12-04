package install

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.AddCommand(installShifuCmd)
}

var installShifuCmd = &cobra.Command{
	Use:     "shifu",
	Aliases: []string{"sf"},
	Short:   "install shifu component in kubernetes cluster",
	Long:    "install shifu component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		yamlURL := "https://raw.githubusercontent.com/Edgenesis/shifu/" + shifu.GetLatestShifuVersion() + "/pkg/k8s/crd/install/shifu_install.yml"
		resp, err := http.Get(yamlURL)
		if err != nil {
			logger.Debug(err)
			logger.Println("Failed to install shifu component")
			return
		}
		defer resp.Body.Close()

		yamlContent, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Debug(err)
			logger.Println("Failed to install shifu component")
			return
		}

		_, err = k8s.ApplyYaml(string(yamlContent))
		if err != nil {
			logger.Debug(err)
			logger.Println("Failed to install shifu component")
			return
		}

		logger.Println("Shifu component installed successfully")
	},
}

func getLatestShifuVersion() string {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "Edgenesis", "shifu", nil)
	if err != nil {
		logger.Println(err)
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

func getAllAvailableVersions() []string {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "Edgenesis", "shifu", nil)
	if err != nil {
		logger.Println(err)
		return nil
	}

	var version []string

	for _, release := range releases {
		version = append(version, *release.Name)
	}

	return version
}
