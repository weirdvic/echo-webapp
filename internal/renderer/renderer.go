package renderer

import (
	"errors"
	"io"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

// Original code from https://machiel.me/post/pongo2-with-echo-or-net-http/

type Renderer struct {
	Debug bool
}

func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var ctx pongo2.Context
	if data != nil {
		var ok bool
		ctx, ok = data.(pongo2.Context)

		if !ok {
			return errors.New("no pongo2.Context data was passed")
		}
	}

	var t *pongo2.Template
	var err error

	if r.Debug {
		t, err = pongo2.FromFile(name)
	} else {
		t, err = pongo2.FromCache(name)
	}

	if err != nil {
		return err
	}

	return t.ExecuteWriter(ctx, w)
}
