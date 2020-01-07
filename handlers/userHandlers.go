package handlers

import (
	user "cmonney/go-web-api/users"
	"net/http"
)

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}
