package template

import (
	"html/template"
	"log"
	"net/http"
)

const UninstalledTemplate = `
<p>Please install shifu first</p>
<button onclick="window.location.href='/install/shifu'">Install shifu</button>
`

func RenderUninstalledTemplate(w http.ResponseWriter) {
	tmpl, err := template.New("webpage").Parse(UninstalledTemplate)
	if err != nil {
		log.Println("error to parse template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
