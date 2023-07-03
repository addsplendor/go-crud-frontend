package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type PostList struct {
	app.Compo
	Posts []blogPost
}

func (p *PostList) Render() app.UI {

	return app.Div().Body(
		app.Range(p.Posts).Slice(func(i int) app.UI {
			return renderBlogPost(p.Posts[i])
		}))
}
