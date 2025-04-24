package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/anicse37/TDD_Go"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("Title: Post 1")},
		"hello_world2.md": {Data: []byte("Title: Post 2")},
	}
	i := 0
	for key, value := range fs {
		t.Run(key, func(t *testing.T) {
			posts, err := blogposts.NewPostsFromFS(fs)
			if err != nil {
				t.Fatal(err)
			}
			got := posts[i]
			want := blogposts.Post{Title: string(value.Data)[7:]}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %+v || Want %+v \n", got, want)
			}
			if len(posts) != len(fs) {
				t.Errorf("Got %+v || Want %v \n", len(posts), len((fs)))
			}
		})
		i++
	}
}
