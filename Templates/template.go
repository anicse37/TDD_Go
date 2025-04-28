package template

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, apost Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\nTags: <ul>", apost.Title, apost.Description)
	if err != nil {
		return err
	}
	for _, tag := range apost.Tags {
		_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
	}
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "</ul>")

	if err != nil {
		return err
	}
	return nil
}
