package controllers

import "github.com/thiagonscosta/api-go/src/middlewares"

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}