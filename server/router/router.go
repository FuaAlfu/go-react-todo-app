package router

import (
	"fmt"
	middleware "server/middleware"

	"github.com/gorilla/mux"

)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/task/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/task/delete/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/task/deleteAllTasks", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return r
}
