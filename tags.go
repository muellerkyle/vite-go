package vueglue

import (
	"bytes"
	"html/template"
	"log"
)

func (vg *VueGlue) RenderTags() (template.HTML, error) {
	var tags string
	log.Println(vg)

	if vg.Environment == "development" {
		tags = `
    <script type="module" src="http://localhost:3000/{{ .MainModule }}"></script>
        `
	} else {
		tags = `
	<script type="module" crossorigin src="/{{ .MainModule }}"></script>
	{{ range .Imports }}
	<link rel="modulepreload" href="/{{.}}">
	{{ end }}
	{{ range .CSSModule }}
	<link rel="stylesheet" href="/{{.}}">
	{{ end }}
	`
	}
	tmpl, err := template.New("tags").Parse(tags)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	tmpl.Execute(&buffer, vg)

	return template.HTML(buffer.String()), nil
}
