package s3

import (
	"context"
	"time"

	"github.com/rai-project/store"
)

const (
	metadataKey      = "github.com/rai-project/store/s3/metadata"
	lifetimeKey      = "github.com/rai-project/store/s3/lifetimeKey"
	fileSizeLimitKey = "github.com/rai-project/store/s3/fileSizeLimitKey"
	aclKey           = "github.com/rai-project/store/s3/acl"
	mimetypeKey      = "github.com/rai-project/store/s3/mimetype"
)

func Metadata(m map[string]*string) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, metadataKey, m)
	}
}

func Lifetime(t time.Duration) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, lifetimeKey, t)
	}
}

func FileSizeLimit(i int64) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, fileSizeLimitKey, i)
	}
}

func MimeType(s string) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, mimetypeKey, s)
	}
}

func ACL(acl string) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, aclKey, acl)
	}
}
