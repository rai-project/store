package s3

import (
	"context"
	"fmt"
	"time"

	"reflect"

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

func toStringPtr(m interface{}) *string {
	// switch m := m.(type) {
	// case bool, int64, int32, int16, int8, int, float64, float32:
	// 	return toStringPtr(fmt.Sprint("%v", m))
	// case string:
	// 	return aws.String(m)
	// default:
	// 	r := reflect.ValueOf(m)
	// 	if r.Kind() == reflect.Ptr {
	// 		r = r.Elem()
	// 	}
	// 	if r.Kind() != reflect.Struct {
	// 		return toStringPtr(fmt.Sprint("%v", m))
	// 	}
	// 	s, err := toMapPtr(structs.Map(m))
	// 	if err != nil {
	// 		return toStringPtr(fmt.Sprint("%v", m))
	// 	}
	// 	return s
	// }
	return aws.String(fmt.Sprint("%v", m))
}

func toMapPtr(m interface{}) (map[string]*string, error) {
	switch m := m.(type) {
	case map[string]*string:
		return m, nil
	case map[string]string:
		out := map[string]*string{}
		for k, v := range m {
			out[k] = toStringPtr(v)
		}
		return out, nil
	case map[string]interface{}:
		out := map[string]*string{}
		for k, v := range m {
			out[k] = toStringPtr(v)
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
