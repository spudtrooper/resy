// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type VenueOption struct {
	f func(*venueOptionImpl)
	s string
}

func (o VenueOption) String() string { return o.s }

type VenueOptions interface {
	DebugBody() bool
	HasDebugBody() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func VenueDebugBody(debugBody bool) VenueOption {
	return VenueOption{func(opts *venueOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}, fmt.Sprintf("api.VenueDebugBody(bool %+v)}", debugBody)}
}
func VenueDebugBodyFlag(debugBody *bool) VenueOption {
	return VenueOption{func(opts *venueOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}, fmt.Sprintf("api.VenueDebugBody(bool %+v)}", debugBody)}
}

func VenueToken(token string) VenueOption {
	return VenueOption{func(opts *venueOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.VenueToken(string %+v)}", token)}
}
func VenueTokenFlag(token *string) VenueOption {
	return VenueOption{func(opts *venueOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.VenueToken(string %+v)}", token)}
}

type venueOptionImpl struct {
	debugBody     bool
	has_debugBody bool
	token         string
	has_token     bool
}

func (v *venueOptionImpl) DebugBody() bool    { return v.debugBody }
func (v *venueOptionImpl) HasDebugBody() bool { return v.has_debugBody }
func (v *venueOptionImpl) Token() string      { return v.token }
func (v *venueOptionImpl) HasToken() bool     { return v.has_token }

type VenueParams struct {
	DebugBody bool   `json:"debug_body"`
	Location  string `json:"location" required:"true"`
	Token     string `json:"token"`
	UrlSlug   string `json:"url_slug" required:"true"`
}

func (o VenueParams) Options() []VenueOption {
	return []VenueOption{
		VenueDebugBody(o.DebugBody),
		VenueToken(o.Token),
	}
}

// ToBaseOptions converts VenueOption to an array of BaseOption
func (o *venueOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
		BaseDebugBody(o.DebugBody()),
	}
}

func makeVenueOptionImpl(opts ...VenueOption) *venueOptionImpl {
	res := &venueOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeVenueOptions(opts ...VenueOption) VenueOptions {
	return makeVenueOptionImpl(opts...)
}
