package template

import (
	"html/template"
	"log"
	"net/http"
)

const devicesHtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Devices</title>
</head>
<body>
    <h1>Devices</h1>
    <ul>
        {{range .}}
        <li><a href="/device/{{.Name}}">{{.Name}}</a></li>
        {{end}}
    </ul>
</body>
</html>
`

func RenderDevicesTemplate(w http.ResponseWriter, links any) {
	tmpl, err := template.New("webpage").Parse(devicesHtmlTemplate)
	if err != nil {
		log.Println("error to parse template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, links)
	if err != nil {
		log.Println("error to execute template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
