package hertz

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

var HeaderXRequestID string

// Option for request id generator
type Option func(*config)

type (
	Generator func() string
	Handler   func(ctx context.Context, c *app.RequestContext, requestID string)
)

// NewRequestID initializes the RequestID middleware.
func NewRequestID(opts ...Option) app.HandlerFunc {
	cfg := &config{
		generator: func() string {
			return uuid.New().String()
		},
		headerKey: "X-Request-ID",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(ctx context.Context, c *app.RequestContext) {
		// Get id from request
		rid := c.Request.Header.Get(string(cfg.headerKey))
		if rid == "" {
			rid = cfg.generator()
		}
		HeaderXRequestID = string(cfg.headerKey)
		if cfg.handler != nil {
			cfg.handler(ctx, c, rid)
		}
		// Set the id to ensure that the request id is in the request and response
		c.Request.Header.Set(HeaderXRequestID, rid)
		c.Response.Header.Set(HeaderXRequestID, rid)
		ctx = context.WithValue(ctx, HeaderXRequestID, rid)
		c.Next(ctx)
	}
}

type HeaderStrKey string

// WithGenerator set generator function
func WithGenerator(g Generator) Option {
	return func(cfg *config) {
		cfg.generator = g
	}
}

// WithCustomHeaderStrKey set custom header key for request id
func WithCustomHeaderStrKey(s HeaderStrKey) Option {
	return func(cfg *config) {
		cfg.headerKey = s
	}
}

// WithHandler set handler function for request id with context
func WithHandler(handler Handler) Option {
	return func(cfg *config) {
		cfg.handler = handler
	}
}

// Config defines the config for RequestID middleware
type config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	generator Generator
	headerKey HeaderStrKey
	handler   Handler
}

// GetRequestID returns the request identifier
func GetRequestID(c *app.RequestContext) string {
	return c.Request.Header.Get(HeaderXRequestID)
}
