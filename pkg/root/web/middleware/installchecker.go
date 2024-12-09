package middleware

import (
	"net/http"

	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/rhoninl/sft/template"
)

func InstallChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := shifu.CheckShifuInstalled(); err != nil {
			template.RenderUninstalledTemplate(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
