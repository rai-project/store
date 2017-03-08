package web

import (
	"github.com/labstack/echo"
	"github.com/rai-project/store"
	"github.com/rai-project/store/s3"
)

func newStore(bucket string) (store.Store, error) {
	var err error
	str, err := s3.New(store.Bucket(bucket))
	if err != nil {
		return nil, err
	}
	return str, err
}

func storeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		store, err := newStore(s3.Config.Bucket)
		if err != nil {
			log.WithError(err).Fatal("Failed to get store context")
			c.Error(err)
		} else {
			c.Set("store", store)
		}
		return next(c)
	}
}

func Init(api *echo.Group) {
	store := api.Group("/store")
	store.Use(storeMiddleware)
	store.GET("/list", list)
	store.POST("/upload", put)
	store.GET("/download/:key", get)
	store.GET("/get/:key", get)
}
