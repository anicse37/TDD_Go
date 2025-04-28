package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/anicse37/TDD_Go"
)

const (
	firstBody = `Title: Post 1
Description: Description 1`
	secondBody = `Title: Post 2
Description: Description 2`
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
			{Title: "Post 1", Description: "Description 1"},
			{Title: "Post 2", Description: "Description 2"},
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
