package s3

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
)

type lifetimeKey struct{}
type expirationKey struct{}
type fileSizeLimitKey struct{}
type aclKey struct{}
type contentTypeKey struct{}
type sessionKey struct{}
type prefixKey struct{}

func NewOptions() *store.Options {
	return &store.Options{
		BaseURL: Config.BaseURL,
		Bucket:  Config.Bucket,
		Context: context.Background(),
	}
}

func toStringPtr(m interface{}) *string {
	switch m := m.(type) {
	case bool, int64, int32, int16, int8, int, float64, float32:
		return toStringPtr(fmt.Sprint("%v", m))
	case string:
		return aws.String(m)
	default:
		buf, err := json.Marshal(m)
		if err != nil {
			return toStringPtr(fmt.Sprint("%v", m))
		}
		return aws.String(string(buf))
	}
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

func Expiration(t time.Time) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, expirationKey{}, t)
	}
}

func Lifetime(t time.Duration) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, lifetimeKey{}, t)
	}
}

func FileSizeLimit(i int64) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, fileSizeLimitKey{}, i)
	}
}

func ContentType(s string) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, contentTypeKey{}, s)
	}
}

func ACL(acl string) store.UploadOption {
	return func(o *store.UploadOptions) {
		o.Context = context.WithValue(o.Context, aclKey{}, acl)
	}
}

func Session(s *session.Session) store.Option {
	return func(o *store.Options) {
		o.Context = context.WithValue(o.Context, sessionKey{}, s)
	}
}

func Prefix(s string) store.ListOption {
	return func(o *store.ListOptions) {
		o.Context = context.WithValue(o.Context, prefixKey{}, s)
	}
}
