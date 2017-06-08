package bolt

import (
	"context"

	"github.com/rai-project/store"
)

const (
	basePathKey = "github.com/rai-project/store/bolt/basePath"
)

func NewOptions() *store.Options {
	return &store.Options{
		Context: context.Background(),
	}
}

func BasePath(s string) store.Option {
	return func(o *store.Options) {
		o.Context = context.WithValue(o.Context, basePathKey, s)
	}
}
