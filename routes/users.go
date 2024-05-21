package routes

import (
	"net/http"
)

// UsersHandler handles requests to the /users endpoint
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic for handling users route
	// For example, you can return a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Users route"}`))
}
