package main

import "log"

func main() {

	Use(func(req Req, res Res, next func()) {
		log.Printf("m1: Request: %s %s\n", req.method, req.path)
		next()
	})

	Use(func(req Req, res Res, next func()) {
		log.Printf("m2: Request: %s %s\n\n", req.method, req.path)
		next()
	})

	Get("/home", func(req Req, res Res) {
		res.send("Hello, World!")
	})

	Post("/home", func(req Req, res Res) {
		reqBody := req.body
		respone := "You sent: " + reqBody
		res.send(respone)
	})

	Get("/post/:postId/comment/:commentId", func(req Req, res Res) {
		postId := req.params["postId"]
		commentId := req.params["commentId"]
		res.send("Post ID: " + postId + ", Comment ID: " + commentId)
	})

	Start(1069)
}
