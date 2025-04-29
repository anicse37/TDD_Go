package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "htmlfiles/*"
	postTemplate embed.FS
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, apost Post) error {
	temp1, err := template.ParseFS(postTemplate, "htmlfiles/*.gohtml")
	if err != nil {
		return err
	}

	if err := temp1.Execute(w, apost); err != nil {
		return err
	}

	return nil
}
