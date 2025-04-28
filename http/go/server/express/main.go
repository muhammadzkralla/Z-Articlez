package main

import "log"

func main() {

	app := App{}

	app.Use(func(req Req, res Res, next func()) {
		log.Printf("m1: Request: %s %s\n", req.method, req.path)
		next()
	})

	app.Use(func(req Req, res Res, next func()) {
		log.Printf("m2: Request: %s %s\n\n", req.method, req.path)
		next()
	})

	app.Get("/home", func(req Req, res Res) {
		res.send("Hello, World!")
	})

	app.Post("/home", func(req Req, res Res) {
		reqBody := req.body
		respone := "You sent: " + reqBody
		res.send(respone)
	})

	app.Get("/post/:postId/comment/:commentId", func(req Req, res Res) {
		postId := req.params["postId"]
		commentId := req.params["commentId"]
		res.send("Post ID: " + postId + ", Comment ID: " + commentId)
	})

	app.Start(1069)
}
