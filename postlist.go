package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type postList struct {
	app.Compo
	posts []blogPost
}

func (p *postList) Render() app.UI {
	var postUIs = []app.UI{}

	for _, post := range p.posts {
		var postUI = renderBlogPost(post)
		postUIs = append(postUIs, postUI)
		app.Logf("Posts: %v\n", post)
	}

	return app.Div().Body(postUIs...)
}