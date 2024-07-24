package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)
func CORS(handler http.Handler) http.Handler{
return handlers.CORS(
	handlers.AllowedOrigins([]string{"http://localhost:5173"}), 
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), 
	handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}), 
)(handler)
}