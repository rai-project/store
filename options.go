package store

import (
	"context"
	"io"
	"strings"

	"github.com/cheggaaa/pb"
)

type Options struct {
	BaseURL string
	Bucket  string
	Context context.Context
}

type Option func(*Options)

func BaseURL(s string) Option {
	return func(o *Options) {
		if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
			s = "http://" + s
		}
		o.BaseURL = s
	}
}

func Bucket(s string) Option {
	return func(o *Options) {
		o.Bucket = s
	}
}

type UploadOptions struct {
	Progress              *pb.ProgressBar
	ProgressOutput        io.Writer
	ProgressFinishMessage string
	Context               context.Context
}

type UploadOption func(*UploadOptions)

func UploadProgress(p *pb.ProgressBar) UploadOption {
	return func(opts *UploadOptions) {
		opts.Progress = p
	}
}

func UploadProgressOutput(out io.Writer) UploadOption {
	return func(opts *UploadOptions) {
		opts.ProgressOutput = out
	}
}

func UploadProgressFinishMessage(s string) UploadOption {
	return func(opts *UploadOptions) {
		opts.ProgressFinishMessage = s
	}
}

type DownloadOptions struct {
	Context context.Context
}

type DownloadOption func(*DownloadOptions)

type GetOptions struct {
	Context context.Context
}

type GetOption func(*GetOptions)

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
