package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rai-project/store"
)

const (
	metadataKey      = "github.com/rai-project/store/s3/metadata"
	lifetimeKey      = "github.com/rai-project/store/s3/lifetime"
	expirationKey    = "github.com/rai-project/store/s3/expiration"
	fileSizeLimitKey = "github.com/rai-project/store/s3/fileSizeLimit"
	aclKey           = "github.com/rai-project/store/s3/acl"
	mimetypeKey      = "github.com/rai-project/store/s3/mimetype"
	sessionKey       = "github.com/rai-project/store/s3/session"
)

func Metadata(m map[string]*string) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, metadataKey, m)
	}
}

func Expiration(t time.Time) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, expirationKey, t)
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

func Session(s *session.Session) store.Option {
	return func(o *store.Options) {
		o.Context = context.WithValue(o.Context, sessionKey, s)
	}
}
