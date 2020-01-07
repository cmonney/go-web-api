package handlers

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

// UsersRouter handles the users route
func UsersRouter(writer http.ResponseWriter, request *http.Request) {

	path := strings.TrimSuffix(request.URL.Path, "/")

	if path == "/users" {
		switch request.Method {
		case http.MethodGet:
			usersGetAll(writer, request)
			return
		case http.MethodPost:
			return
		default:
			postError(writer, http.StatusMethodNotAllowed)
		}
	}

	path = strings.TrimPrefix(path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(writer, http.StatusNotFound)
		return
	}

	// id := bson.ObjectIdHex(path)

	switch request.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(writer, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(request.URL.Path)
}
