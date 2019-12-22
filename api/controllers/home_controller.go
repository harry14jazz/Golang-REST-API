package controllers

import (
	"net/http"

	"github.com/harry14jazz/fullstack/api/responses"
)

//Home API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
