package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"
	"github.com/approvals/go-approval-tests"
	"github.com/isedaniel/blogrenderer"
)

func TestRender(t *testing.T) {
    var (
        post1Body = `# Body Title
This is the body of the post.

A little link [here](github.com/isedaniel).`
    )

	var (
		p = blogrenderer.Post{
			Title:       "Hello World",
            Body: post1Body,
			Description: "This is a description.",
			Tags:        []string{"go", "tdd"},
		}
	)

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("single post to HTML", func(t *testing.T) {
		b := bytes.Buffer{}
		if err := pr.Render(&b, p); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, b.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		p = blogrenderer.Post{
			Title:       "Hello, World!",
			Body:        "This is a post.",
			Description: "This is a description.",
			Tags:        []string{"go", "tdd"},
		}
	)
    pr, err := blogrenderer.NewPostRenderer()
    if err != nil {
        b.Fatal(err)
    }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pr.Render(io.Discard, p)
	}
}
