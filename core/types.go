package core

// Exchange is the envelope that holds In and Out message between endpoints and processors within a route.
type Exchange struct {
	RouteID string
	In      Message
	Out     Message
}

// Message is the unit of data that is being consumed and produced by endpoints and processors within a route.
type Message struct {
	Header map[string][]string
	Body   []byte
}

// Endpoint describes endpoint that can be consumed by different components
type Endpoint interface {
	Process(xchng Exchange)
}

// RouteBuilder is the interface used to build routes using DSL.
type RouteBuilder interface {
	From(e ...Endpoint) RouteBuilder
	To(e ...Endpoint) RouteBuilder
}
