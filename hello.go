package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io"
	"net/http"
	"time"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
	Posts []blogPost
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {

	app.Logf("%d Posts: %v\n", len(h.Posts), h.Posts)

	return app.Div().Body(
		&PostList{
			Posts: h.Posts,
		},
		app.Button().Text("Hello World!").OnClick(func(ctx app.Context, e app.Event) {
			//h.Posts = append(h.Posts, blogPost{
			//	title:   "The redundant array",
			//	author:  "Addie Daria",
			//	content: "Junior baby programmer learn how write good",
			//	date:    time.Now(),
			//})

			ctx.Async(func() {
				r, err := http.Get("http://localhost:8080/posts")
				if err != nil {
					app.Log(err)
					return
				}
				defer r.Body.Close()

				b, err := io.ReadAll(r.Body)
				if err != nil {
					app.Log(err)
					return
				}

				app.Logf("request response: %s", b)
			})
		}),

		&submit{

		},
	)

	//return app.Div().Body(
	//	app.Div(),
	//	app.Textarea().Text(h.text),
	//	app.Button().Text(\"Hello World!").OnClick(func(ctx app.Context, e app.Event) {
	//		h.text = "Hello World!"
	//		h.Update()
	//
	//	})

}

type blogPost struct {
	title   string
	author  string
	content string
	date    time.Time
}

func renderBlogPost(blogPost blogPost) app.UI {
	return app.Div().Body(
		app.H1().Text(blogPost.title),
		app.H3().Text(fmt.Sprintf("%s %s", blogPost.author, fmt.Sprintf("%d-%d-%d %d:%02d\n",
			blogPost.date.Year(),
			blogPost.date.Month(),
			blogPost.date.Day(),
			blogPost.date.Hour(),
			blogPost.date.Minute()),
		)),
		app.P().Text(blogPost.content).Style("background-color", "blue"),
	)
}

func (h *hello) OnMount(ctx app.Context) {
	// temporarily hardcode the data!
	h.Posts = []blogPost{
		blogPost{
			title:   "The Downfall of Neo-Noir",
			author:  "Addie Daria",
			content: "Pretentious drivel",
			date:    time.Date(2023, time.May, 25, 18, 50, 00, 05, time.Local),
		},

		blogPost{
			title:   "Why Carrots are Problematic",
			author:  "Addie Daria",
			content: "AGGGGGGGGHHHHHHHHHHH",
			date:    time.Date(2023, time.May, 31, 12, 26, 00, 00, time.Local),
		},
	}
	h.Update()
}
