package controllers

import (
	"github.com/gorilla/mux"

	"github.com/bassman7689/fetch-exercise/pkg/store"
)

func Register(r *mux.Router, store store.Store) {
	receiptsController := &ReceiptsController{store: store}
	r.HandleFunc("/receipts/process", receiptsController.Process).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", receiptsController.Points).Methods("GET")
}
