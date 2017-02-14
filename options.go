package store

import "context"

type Options struct {
	BaseURL string
	Bucket  string
	Context context.Context
}

type Option func(*Options)

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
