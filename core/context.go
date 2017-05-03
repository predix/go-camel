package core

// Context represents a single camel routing rule base. It is the starting point
type Context interface {
	// Lifecycle methods
	Start() error
	Stop() error
	Suspend() error
	Resume() error

	AddRoute(routeBuilder RouteBuilder) error
}
