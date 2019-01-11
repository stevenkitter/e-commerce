package server

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type VideoApi struct {
	DB gorm.DB
}

func (api *VideoApi) Endpoint(r *mux.Router) {
	router := r.PathPrefix("/api").Subrouter()
	router.HandleFunc(
		"/course",
		WrapRestful(
			api.PostCourse,
			api.DeleteCourse,
			api.PutCourse,
			api.GetCourse))
}
