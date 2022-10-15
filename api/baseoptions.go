// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type BaseOption func(*baseOptionImpl)

type BaseOptions interface {
	Token() string
	HasToken() bool
	DebugBody() bool
	HasDebugBody() bool
}

func BaseToken(token string) BaseOption {
	return func(opts *baseOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func BaseTokenFlag(token *string) BaseOption {
	return func(opts *baseOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func BaseDebugBody(debugBody bool) BaseOption {
	return func(opts *baseOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func BaseDebugBodyFlag(debugBody *bool) BaseOption {
	return func(opts *baseOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

type baseOptionImpl struct {
	token         string
	has_token     bool
	debugBody     bool
	has_debugBody bool
}

func (b *baseOptionImpl) Token() string      { return b.token }
func (b *baseOptionImpl) HasToken() bool     { return b.has_token }
func (b *baseOptionImpl) DebugBody() bool    { return b.debugBody }
func (b *baseOptionImpl) HasDebugBody() bool { return b.has_debugBody }

func makeBaseOptionImpl(opts ...BaseOption) *baseOptionImpl {
	res := &baseOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeBaseOptions(opts ...BaseOption) BaseOptions {
	return makeBaseOptionImpl(opts...)
}
