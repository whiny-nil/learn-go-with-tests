package blog_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"testing/fstest"

	approvals "github.com/approvals/go-approval-tests"
	blog "github.com/whiny-nil/learn-go-with-tests/18-templating/blog"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: foo, bar
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blog.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blog.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello
World`}

	assertPost(t, got, want)
}

func TestRender(t *testing.T) {
	var (
		aPost = blog.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blog.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blog.Post{{Title: "hello world"}, {Title: "hello world 2"}}

		err := postRenderer.RenderIndex(&buf, posts)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blog.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blog.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}

func assertPost(t *testing.T, got blog.Post, want blog.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
