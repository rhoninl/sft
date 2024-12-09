package devices

import (
	"net/http"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/template"
)

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	device, err := k8s.GetAllByDeviceName(r.PathValue("device_name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template.RenderDetailTemplate(w, device)
}
