package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/anicse37/TDD_Go"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, docker
Body: Hello
Aniket`

	secondBody = `Title: Post 2
Description: Description 2
Tags: tdd, go
Body: Hello
This is gtech`
)

func TestNewBlogPosts(t *testing.T) {
	var i int
	for i < 2 {

		fs := fstest.MapFS{
			"hello-world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}
		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
		// rest of test code cut for brevity
		want := []blogposts.Post{
			{
				Title:       "Post 1",
				Description: "Description 1",
				Tags: []string{
					"tdd",
					"docker",
				},
				Body: `Hello
Aniket
`},
			{
				Title:       "Post 2",
				Description: "Description 2",
				Tags: []string{
					"tdd",
					"go",
				},
				Body: `Hello
This is gtech
`},
		}

		assertPost(t, posts[i], want[i])
		i++
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
