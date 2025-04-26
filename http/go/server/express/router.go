package main

import "strings"

var getRoutes []Route
var postRoutes []Route

func Get(endpoint string, handler Handler) {
	getRoutes = append(getRoutes, Route{endpoint, applyMiddleware(handler)})
}

func Post(endpoint string, handler Handler) {
	postRoutes = append(postRoutes, Route{endpoint, applyMiddleware(handler)})
}

func matchRoute(requestPath string, routes []Route) (Handler, map[string]string) {
	for _, route := range routes {
		params := make(map[string]string)
		routeParts := strings.Split(route.path, "/")
		requestParts := strings.Split(requestPath, "/")

		if len(routeParts) != len(requestParts) {
			continue
		}

		match := true
		for i := range routeParts {
			if strings.HasPrefix(routeParts[i], ":") {
				paramName := routeParts[i][1:]
				params[paramName] = requestParts[i]
			} else if routeParts[i] != requestParts[i] {
				match = false
				break
			}
		}

		if match {
			return route.handler, params
		}
	}

	return nil, nil
}
