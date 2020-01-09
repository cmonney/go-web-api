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
			usersPostOne(writer, request)
			return
		case http.MethodHead:
			usersGetAll(writer, request)
			return
		case http.MethodOptions:
			postOptionsResponse(writer, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(writer, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(writer, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)

	switch request.Method {
	case http.MethodGet:
		usersGetOne(writer, request, id)
		return
	case http.MethodPut:
		usersPutOne(writer, request, id)
		return
	case http.MethodPatch:
		usersPatchOne(writer, request, id)
		return
	case http.MethodDelete:
		usersDeleteOne(writer, request, id)
		return
	case http.MethodHead:
		usersGetOne(writer, request, id)
		return
	case http.MethodOptions:
		postOptionsResponse(writer, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(writer, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(request.URL.Path)
}
