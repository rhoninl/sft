package shifu

import (
	"encoding/json"
	"net/http"

	"github.com/rhoninl/sft/pkg/utils/shifu"
)

type InstallCheckResponse struct {
	Installed bool `json:"installed"`
}

func InstallChecker(w http.ResponseWriter, r *http.Request) {
	response := InstallCheckResponse{
		Installed: true,
	}

	if err := shifu.CheckShifuInstalled(); err != nil {
		response.Installed = false
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
