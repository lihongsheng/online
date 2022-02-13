// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetYakitPluginURL generates an URL for the get yakit plugin operation
type GetYakitPluginURL struct {
	Limit   *int64
	Name    *string
	Order   *string
	OrderBy *string
	Page    *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetYakitPluginURL) WithBasePath(bp string) *GetYakitPluginURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetYakitPluginURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetYakitPluginURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/yakit/plugin"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var nameQ string
	if o.Name != nil {
		nameQ = *o.Name
	}
	if nameQ != "" {
		qs.Set("name", nameQ)
	}

	var orderQ string
	if o.Order != nil {
		orderQ = *o.Order
	}
	if orderQ != "" {
		qs.Set("order", orderQ)
	}

	var orderByQ string
	if o.OrderBy != nil {
		orderByQ = *o.OrderBy
	}
	if orderByQ != "" {
		qs.Set("order_by", orderByQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt64(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetYakitPluginURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetYakitPluginURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetYakitPluginURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetYakitPluginURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetYakitPluginURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetYakitPluginURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}