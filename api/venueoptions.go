// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type VenueOption func(*venueOptionImpl)

type VenueOptions interface {
	Token() string
	HasToken() bool
	DebugBody() bool
	HasDebugBody() bool
	ToBaseOptions() []BaseOption
}

func VenueToken(token string) VenueOption {
	return func(opts *venueOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func VenueTokenFlag(token *string) VenueOption {
	return func(opts *venueOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func VenueDebugBody(debugBody bool) VenueOption {
	return func(opts *venueOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func VenueDebugBodyFlag(debugBody *bool) VenueOption {
	return func(opts *venueOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

type venueOptionImpl struct {
	token         string
	has_token     bool
	debugBody     bool
	has_debugBody bool
}

func (v *venueOptionImpl) Token() string      { return v.token }
func (v *venueOptionImpl) HasToken() bool     { return v.has_token }
func (v *venueOptionImpl) DebugBody() bool    { return v.debugBody }
func (v *venueOptionImpl) HasDebugBody() bool { return v.has_debugBody }

type VenueParams struct {
	UrlSlug   string `json:"url_slug" required:"true"`
	Location  string `json:"location" required:"true"`
	Token     string `json:"token"`
	DebugBody bool   `json:"debug_body"`
}

func (o VenueParams) Options() []VenueOption {
	return []VenueOption{
		VenueToken(o.Token),
		VenueDebugBody(o.DebugBody),
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
		opt(res)
	}
	return res
}

func MakeVenueOptions(opts ...VenueOption) VenueOptions {
	return makeVenueOptionImpl(opts...)
}
