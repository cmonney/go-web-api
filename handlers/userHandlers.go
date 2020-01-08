package handlers

import (
	user "cmonney/go-web-api/users"
	"errors"
	"net/http"
)

func bodyToUser(r *http.Request, u *user.User) error {
	if r.Body == nil {
		return errors.New("Request body is empty")
	}
}

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}
