package controllers

import (
	"net/http"
	"online-store-evermos/src/responses"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome Online Store API Evermos")
}
