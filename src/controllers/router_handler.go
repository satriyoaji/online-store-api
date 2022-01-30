package controllers

import "online-store-evermos/src/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Items Routes
	s.Router.HandleFunc("/item", middlewares.SetMiddlewareJSON(s.CreateItem)).Methods("POST")
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.GetItems)).Methods("GET")
	s.Router.HandleFunc("/item/{id}", middlewares.SetMiddlewareJSON(s.GetItem)).Methods("GET")
	s.Router.HandleFunc("/item/{id}", middlewares.SetMiddlewareJSON(s.UpdateItem)).Methods("PUT")
}
