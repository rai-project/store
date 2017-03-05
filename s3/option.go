package s3

import (
	"context"
	"time"

	"reflect"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
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

func toMapPtr(m interface{}) (map[string]*string, error) {
	switch m.(type) {
	case map[string]*string:
		return m.(map[string]*string), nil
	case map[string]string:
		out := map[string]*string{}
		for k, v := range m.(map[string]string) {
			out[k] = aws.String(v)
		}
		return out, nil
	case map[string]interface{}:
		out := map[string]*string{}
		for k, v := range m.(map[string]interface{}) {
			g := fmt.Sprint(v)
			out[k] = aws.String(g)
		}
		return out, nil
	default:
		r := reflect.ValueOf(m)
		if r.Kind() == reflect.Ptr {
			r = r.Elem()
		}
		if r.Kind() != reflect.Struct {
			return map[string]*string{}, errors.Errorf("input %v must be a map or a struct", m)
		}
		return toMapPtr(structs.Map(m))
	}
}

func Metadata(m interface{}) store.UploadOption {
	return func(o *store.UploadOptions) {
		v, err := toMapPtr(m)
		if err != nil {
			log.WithError(err).Error("invalid s3 metadata")
			return
		}
		o.Context = context.WithValue(o.Context, metadataKey, v)
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
