package template

import (
	"html/template"
	"log"
	"net/http"

	"github.com/rhoninl/sft/pkg/k8s"
)

const detailHtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.EdgeDevice.Name}}</title>
</head>
<body>
    <h1><a href="/device"><-Back</a> {{.EdgeDevice.Name}}</h1>
    <ul>
		<li>Name: {{.EdgeDevice.Name}}</li>
    </ul>
</body>
</html>
`

func RenderDetailTemplate(w http.ResponseWriter, deviceInfo *k8s.Device) {
	tmpl, err := template.New("detail").Parse(detailHtmlTemplate)
	if err != nil {
		log.Println("error to parse template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, deviceInfo)
	if err != nil {
		log.Println("error to execute template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
