package web

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/device", http.StatusPermanentRedirect)
}