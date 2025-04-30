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

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "htmlfiles/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderBottom(w io.Writer) error {
	return r.templ.ExecuteTemplate(w, "bottom", nil)
}
func (r *PostRenderer) RenderBlog(w io.Writer) error {
	return r.templ.ExecuteTemplate(w, "blog", nil)
}
func (r *PostRenderer) RenderTop(w io.Writer) error {
	return r.templ.ExecuteTemplate(w, "top", nil)
}

func Render(w io.Writer, apost Post) error {
	temp1, err := template.ParseFS(postTemplate, "htmlfiles/*.gohtml")
	if err != nil {
		return err
	}
	if err := temp1.ExecuteTemplate(w, "blog.gohtml", apost); err != nil {
		return err
	}

	return nil
}
