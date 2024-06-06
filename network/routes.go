package network

type Handler func(request Request, response *Response)

type Route struct {
	Method  HTTPMethod
	Path    Path
	Handler Handler
}

type RoutesMap map[HTTPMethod]map[Path]*Route

func (routesMap *RoutesMap) AddHandler(method HTTPMethod, path Path, handler Handler) {

	// Initialize the inner map if it's nil
	if (*routesMap) == nil {
		*routesMap = make(RoutesMap)
	}

	if (*routesMap)[method] == nil {
		(*routesMap)[method] = make(map[Path]*Route)
	}
	(*routesMap)[method][path] = &Route{Method: method, Path: path, Handler: handler}
}

func (routesMap *RoutesMap) Get(method HTTPMethod, path Path) *Route {
	return (*routesMap)[method][path]
}
