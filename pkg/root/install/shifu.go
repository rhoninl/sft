package install

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/rhoninl/sft/pkg/k8s"
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
		yamlURL := "https://gitee.com/edgenesis/shifu/raw/" + getLatestShifuVersion() + "/pkg/k8s/crd/install/shifu_install.yml"
		resp, err := http.Get(yamlURL)
		if err != nil {
			fmt.Println("Failed to install shifu component")
			return
		}
		defer resp.Body.Close()

		yamlContent, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to install shifu component")
			return
		}

		_, err = k8s.ApplyYaml(string(yamlContent))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to install shifu component")
			return
		}

		fmt.Println("Shifu component installed successfully")
	},
}

func getLatestShifuVersion() string {
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

func getAllAvailableVersions() []string {
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
