package main

var middlewares []Middleware

func Use(m Middleware) {
	middlewares = append(middlewares, m)
}

func applyMiddleware(finalHandler Handler) Handler {
	return func(req Req, res Res) {
		i := 0

		var next func()
		next = func() {
			if i < len(middlewares) {
				middleware := middlewares[i]
				i++
				middleware(req, res, next)
			} else {
				finalHandler(req, res)
			}
		}

		next()
	}
}
