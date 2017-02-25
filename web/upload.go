package web

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/rai-project/store"
	"github.com/rai-project/uuid"
)

func put(c echo.Context) error {
	store := c.Get("store").(store.Store)
	file, err := c.FormFile("files")
	if err != nil {
		pp.Println("ERRROR..... " + err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	defer src.Close()

	id := uuid.NewV4()
	if err := store.UploadFrom(src, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"id": id,
	})
}
