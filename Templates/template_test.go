package template_test

import (
	"bytes"
	"testing"

	template "github.com/anicse37/TDD_Go/Templates"
)

func TestTemplateing(t *testing.T) {
	var (
		aPost = template.Post{
			Title:       "Title of Index",
			Description: "This is the Description of this website's first page",
			Tags:        []string{"website", "tdd"},
			Body:        "This is the bosy of the website, for it contains all the information",
		}
	)

	t.Run("For Single functions", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := template.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Title of Index</h1>
<p>This is the Description of this website's first page</p>
Tags: <ul><li>website</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
