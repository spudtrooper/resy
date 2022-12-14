// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type ConfigOption struct {
	f func(*configOptionImpl)
	s string
}

func (o ConfigOption) String() string { return o.s }

type ConfigOptions interface {
	DebugBody() bool
	HasDebugBody() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func ConfigDebugBody(debugBody bool) ConfigOption {
	return ConfigOption{func(opts *configOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}, fmt.Sprintf("api.ConfigDebugBody(bool %+v)", debugBody)}
}
func ConfigDebugBodyFlag(debugBody *bool) ConfigOption {
	return ConfigOption{func(opts *configOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}, fmt.Sprintf("api.ConfigDebugBody(bool %+v)", debugBody)}
}

func ConfigToken(token string) ConfigOption {
	return ConfigOption{func(opts *configOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.ConfigToken(string %+v)", token)}
}
func ConfigTokenFlag(token *string) ConfigOption {
	return ConfigOption{func(opts *configOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.ConfigToken(string %+v)", token)}
}

type configOptionImpl struct {
	debugBody     bool
	has_debugBody bool
	token         string
	has_token     bool
}

func (c *configOptionImpl) DebugBody() bool    { return c.debugBody }
func (c *configOptionImpl) HasDebugBody() bool { return c.has_debugBody }
func (c *configOptionImpl) Token() string      { return c.token }
func (c *configOptionImpl) HasToken() bool     { return c.has_token }

type ConfigParams struct {
	DebugBody bool   `json:"debug_body"`
	Token     string `json:"token"`
	VenueID   int    `json:"venue_id" required:"true"`
}

func (o ConfigParams) Options() []ConfigOption {
	return []ConfigOption{
		ConfigDebugBody(o.DebugBody),
		ConfigToken(o.Token),
	}
}

// ToBaseOptions converts ConfigOption to an array of BaseOption
func (o *configOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseDebugBody(o.DebugBody()),
		BaseToken(o.Token()),
	}
}

func makeConfigOptionImpl(opts ...ConfigOption) *configOptionImpl {
	res := &configOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeConfigOptions(opts ...ConfigOption) ConfigOptions {
	return makeConfigOptionImpl(opts...)
}
