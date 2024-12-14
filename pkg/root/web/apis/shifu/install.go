package shifu

import (
	"net/http"

	"github.com/rhoninl/sft/pkg/root/install"
)

func InstallShifu(w http.ResponseWriter, r *http.Request) {
	version := r.URL.Query().Get("version")
	if version == "" {
		version = install.EmptyVersion
	}

	err := install.InstallShifu(version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
