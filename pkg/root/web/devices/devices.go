package devices

import (
	"net/http"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/template"
)

func DevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := k8s.GetEdgedevices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template.RenderDevicesTemplate(w, devices)
}
