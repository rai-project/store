package store

import (
	"context"

	"github.com/rai-project/store"
)

type Options struct {
	BaseURL string
	Bucket  string
	Context context.Context
}

type Option func(*Options)

func BaseURL(s string) store.Option {
	return func(o *store.UploadOptions) {
		o.BaseURL = s
	}
}

func Bucket(s string) store.Option {
	return func(o *store.UploadOptions) {
		o.Bucket = s
	}
}

type UploadOptions struct {
	Context context.Context
}

type UploadOption func(*UploadOptions)

type DownloadOptions struct {
	Context context.Context
}

type DownloadOption func(*DownloadOptions)

type ListOptions struct {
	Max     int64
	Context context.Context
}

type ListOption func(*ListOptions)

func Max(m int64) store.ListOption {
	return func(o *store.UploadOptions) {
		o.Max = m
	}
}
