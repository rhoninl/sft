package install

import (
	"net/http"

	"github.com/rhoninl/sft/pkg/root/install"
)

func InstallHandler(w http.ResponseWriter, r *http.Request) {
	install.InstallShifu("")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
