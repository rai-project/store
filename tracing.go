package store

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/opentracing template

//go:generate gowrap gen -d . -i Store -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/opentracing -o tracing.go

// StoreWithTracing implements Store interface instrumented with opentracing spans
type StoreWithTracing struct {
	Store
	_instance string
}

// NewStoreWithTracing returns StoreWithTracing
func NewStoreWithTracing(base Store, instance string) StoreWithTracing {
	return StoreWithTracing{
		Store:     base,
		_instance: instance,
	}
}
