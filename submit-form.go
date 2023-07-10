package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type submit struct {
	app.Compo
	title   string
	author  string
	content string
}

func (s *submit) Render() app.UI {

	return app.Div().Body(
			app.Input().
				Type("text").
				Value(s.title).
				Placeholder("What is the title of your blog post?").
				AutoFocus(true).
				OnChange(s.ValueTo(&s.title)).Style("display", "block"),

			app.Input().
				Type("text").
				Value(s.author).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(s.ValueTo(&s.author)).Style("display", "block"),

			app.Textarea().
				Text(s.content).
				Placeholder("Enter your content here.").
				AutoFocus(true).
				OnChange(s.ValueTo(&s.content)).Style("display", "block"),

			app.Button().Text("Submit your post!").OnClick(func(ctx app.Context, e app.Event){}),

	)}