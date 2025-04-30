package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	blogrenderer "github.com/anicse37/TDD_Go/Templates"
	approvals "github.com/approvals/go-approval-tests"
)

/*-------------------------------------------------------------*/
var (
	Post1 = blogrenderer.Post{
		Title:       "Title of Index",
		Description: "This is the Description of this websites first page",
		Tags:        []string{"website", "tdd"},
		Body:        "This is the bosy of the website, for it contains all the information",
	}
)

/*-------------------------------------------------------------*/

func TestSimpleFunction(t *testing.T) {
	buf := bytes.Buffer{}
	err := blogrenderer.Render(&buf, Post1)

	if err != nil {
		t.Fatal(err)
	}

	got := buf.String()
	want := `<h1>Title of Index</h1>
<p>This is the Description of this websites first page</p>
Tags: <ul><li>website</li><li>tdd</li></ul>
`
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

/*--------------------------------------------------------------*/
func TestApprovals(t *testing.T) {
	buf := bytes.Buffer{}
	if err := blogrenderer.Render(&buf, Post1); err != nil {
		t.Fatal(err)
	}
	approvals.VerifyString(t, buf.String())
}

/*-------------------------------------------------------------*/
func TestRenderBlog(t *testing.T) {
	buf := bytes.Buffer{}
	if err := blogrenderer.Render(&buf, Post1); err != nil {
		t.Fatal(err)
	}
	approvals.VerifyString(t, buf.String())

}

/*-------------------------------------------------------------*/

func TestRenderBottom(t *testing.T) {
	pr, _ := blogrenderer.NewPostRenderer()
	var buf bytes.Buffer
	err := pr.RenderBottom(&buf)
	if err != nil {
		t.Fatal(err)
	}
	approvals.VerifyString(t, buf.String())
}

/*-------------------------------------------------------------*/
func TestRenderTop(t *testing.T) {
	pr, _ := blogrenderer.NewPostRenderer()
	var buf bytes.Buffer
	err := pr.RenderTop(&buf)
	if err != nil {
		t.Fatal(err)
	}
	approvals.VerifyString(t, buf.String())
}

/*-------------------------------------------------------------*/
func BenchmarkRender(b *testing.B) {
	var (
		Post1 = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, Post1)
	}

}

/*-------------------------------------------------------------*/
