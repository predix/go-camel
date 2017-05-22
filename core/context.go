package core

import (
	c "context"
)

// Context represents a single camel routing rule base. It is the starting point
type Context interface {
	// Lifecycle methods
	Start() error
	Stop() error
	Suspend() error
	Resume() error

	AddRoute(routeBuilder RouteBuilder) error
}

type camelContext struct {
	c.Context
}

func (c *camelContext) AddRoute(routeBuilder RouteBuilder) error {
	return nil
}

func (c *camelContext) Resume() error {
	return nil
}

func (c *camelContext) Suspend() error {
	return nil
}

func (c *camelContext) Stop() error {
	return nil
}

func (c *camelContext) Start() error {
	return nil
}

func NewContext() Context {
	return &camelContext{}
}
