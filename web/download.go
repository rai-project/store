package web

import (
	"net/http"

	"github.com/labstack/echo"
)

func get(c echo.Context) error {
	store := c.Get("store").(*tore.Store)
	bytes, err := store.Get(c.Param("key"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	contentType := http.DetectContentType(bytes)
	return c.Blob(http.StatusOK, contentType, bytes)
}
