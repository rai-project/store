package web

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo"
	"github.com/rai-project/store"
)

func get(c echo.Context) error {
	store := c.Get("store").(store.Store)
	buf := new(aws.WriteAtBuffer)
	err := store.DownloadTo(buf, c.Param("key"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	bytes := buf.Bytes()
	contentType := http.DetectContentType(bytes)
	return c.Blob(http.StatusOK, contentType, bytes)
}
