package engine

import "github.com/predix/go-camel/core"

type routeBuilder struct {
}

func (r *routeBuilder) From(core ...core.Endpoint) core.RouteBuilder {
	return r
}

func (r *routeBuilder) To(core ...core.Endpoint) core.RouteBuilder {
	return r
}

func NewRouteBuilder() core.RouteBuilder {
	return &routeBuilder{}
}
