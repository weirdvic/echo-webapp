package main

import (
	"net/http"
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

//e.GET("/", indexPage)
func (app *application) indexPage(c echo.Context) error {
	s, err := app.snippets.Latest()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.Render(http.StatusOK, "ui/html/index.html", pongo2.Context{"snippets": s})
}

//e.GET("/snippet", showSnippet)
func (app *application) showSnippet(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil || id < 1 {
		return c.String(http.StatusNotFound, "Snippet not found!")
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Snippet not found!")
	} else {
		//return c.JSON(http.StatusOK, s)
		return c.Render(http.StatusOK, "ui/html/snippet.html", pongo2.Context{"snippet": s})
	}
}

//e.GET("/snippet/create", createSnippet)
func (app *application) createSnippet(c echo.Context) error {
	return c.String(http.StatusOK, "Creating the snippet")
}
