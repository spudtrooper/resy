// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type ConfigOption func(*configOptionImpl)

type ConfigOptions interface {
	Token() string
	HasToken() bool
	DebugBody() bool
	HasDebugBody() bool
	DebugPayload() bool
	HasDebugPayload() bool
	ToBaseOptions() []BaseOption
}

func ConfigToken(token string) ConfigOption {
	return func(opts *configOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func ConfigTokenFlag(token *string) ConfigOption {
	return func(opts *configOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func ConfigDebugBody(debugBody bool) ConfigOption {
	return func(opts *configOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func ConfigDebugBodyFlag(debugBody *bool) ConfigOption {
	return func(opts *configOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

func ConfigDebugPayload(debugPayload bool) ConfigOption {
	return func(opts *configOptionImpl) {
		opts.has_debugPayload = true
		opts.debugPayload = debugPayload
	}
}
func ConfigDebugPayloadFlag(debugPayload *bool) ConfigOption {
	return func(opts *configOptionImpl) {
		if debugPayload == nil {
			return
		}
		opts.has_debugPayload = true
		opts.debugPayload = *debugPayload
	}
}

type configOptionImpl struct {
	token            string
	has_token        bool
	debugBody        bool
	has_debugBody    bool
	debugPayload     bool
	has_debugPayload bool
}

func (c *configOptionImpl) Token() string         { return c.token }
func (c *configOptionImpl) HasToken() bool        { return c.has_token }
func (c *configOptionImpl) DebugBody() bool       { return c.debugBody }
func (c *configOptionImpl) HasDebugBody() bool    { return c.has_debugBody }
func (c *configOptionImpl) DebugPayload() bool    { return c.debugPayload }
func (c *configOptionImpl) HasDebugPayload() bool { return c.has_debugPayload }

type ConfigParams struct {
	VenueID      int    `json:"venue_id" required:"true"`
	Token        string `json:"token"`
	DebugBody    bool   `json:"debug_body"`
	DebugPayload bool   `json:"debug_payload"`
}

func (o ConfigParams) Options() []ConfigOption {
	return []ConfigOption{
		ConfigToken(o.Token),
		ConfigDebugBody(o.DebugBody),
		ConfigDebugPayload(o.DebugPayload),
	}
}

// ToBaseOptions converts ConfigOption to an array of BaseOption
func (o *configOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
		BaseDebugBody(o.DebugBody()),
		BaseDebugPayload(o.DebugPayload()),
	}
}

func makeConfigOptionImpl(opts ...ConfigOption) *configOptionImpl {
	res := &configOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeConfigOptions(opts ...ConfigOption) ConfigOptions {
	return makeConfigOptionImpl(opts...)
}
