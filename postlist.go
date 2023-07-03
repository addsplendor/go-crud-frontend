package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type postList struct {
	app.Compo
	posts []blogPost
}

func (p *postList) Render() app.UI {

	return app.Div().Body(
		app.Range(p.posts).Slice(func(i int) app.UI {
			return renderBlogPost(p.posts[i])
		}))
}
