package web

import (
	"net/http"

	store "bitbucket.org/c3sr/p3sr-store"
	"github.com/labstack/echo"
)

func list(c echo.Context) error {
	store := c.Get("store").(store.Store)
	keys, err := store.List()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, keys)
}
