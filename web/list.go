package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rai-project/store"
)

func list(c echo.Context) error {
	store := c.Get("store").(store.Store)
	keys, err := store.List()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, keys)
}
