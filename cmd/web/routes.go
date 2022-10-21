package main

import r "github.com/weirdvic/echo-tutorial/internal/renderer"

func (app *application) init() {
	// Set up routing
	app.echo.Static("/", "ui/static")
	app.echo.GET("/", app.indexPage)
	app.echo.GET("/snippet", app.showSnippet)
	app.echo.POST("/snippet/create", app.createSnippet)
	// Attach template renderer
	app.echo.Renderer = r.Renderer{
		Debug: false,
	}
}
