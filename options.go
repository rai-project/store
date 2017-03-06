package store

import (
	"context"
)

type Options struct {
	BaseURL string
	Bucket  string
	Context context.Context
}

type Option func(*Options)

func BaseURL(s string) Option {
	return func(o *Options) {
		o.BaseURL = s
	}
}

func Bucket(s string) Option {
	return func(o *Options) {
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

func Max(m int64) ListOption {
	return func(o *ListOptions) {
		o.Max = m
	}
}

type DeleteOption func(*ListOptions)

type DeleteOptions struct {
	Context context.Context
}
