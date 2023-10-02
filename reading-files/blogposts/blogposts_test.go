package blogposts_test

import (
	"errors"
	"github.com/isedaniel/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("2 files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, but wanted %+v", got, want)
		}
	})

	t.Run("error handling", func(t *testing.T) {
		fs := failingFS{}

		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})
}

type failingFS struct{}

func (fs failingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Sry, I'll allways fail.")
}