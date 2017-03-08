package web

import (
	"github.com/labstack/echo"
	"github.com/rai-project/store"
	"github.com/rai-project/store/s3"
)

func storeMiddleware(opts ...store.Option) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store, err := s3.New(opts...)
			if err != nil {
				log.WithError(err).Fatal("Failed to get store context")
				c.Error(err)
			} else {
				c.Set("store", store)
			}
			return next(c)
		}
	}
}

func Init(api *echo.Group, opts ...store.Option) {
	store := api.Group("/store")
	store.Use(storeMiddleware(opts...))
	store.GET("/list", list)
	store.POST("/upload", put)
	store.GET("/download/:key", get)
	store.GET("/get/:key", get)
}
