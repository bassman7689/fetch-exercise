package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/bassman7689/fetch-exercise/pkg/requests"
	"github.com/bassman7689/fetch-exercise/pkg/responses"
	"github.com/bassman7689/fetch-exercise/pkg/store"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type ReceiptsController struct {
	store store.Store
}

func handleResponse(w http.ResponseWriter, statusCode int, res any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}


func handleError(w http.ResponseWriter, statusCode int, message string) {
	handleResponse(w, statusCode, &responses.APIError{Error: message})
}

func (rc *ReceiptsController) Process(w http.ResponseWriter, r *http.Request) {
	prr := &requests.ProcessReceipt{}
	if err := json.NewDecoder(r.Body).Decode(prr); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid json in request body")
		return
	}

	if err := validate.Struct(prr); err != nil {
		// TODO: better error response
		handleError(w, http.StatusBadRequest, "Invalid fields in request json")
		return
	}

	id, err := rc.store.ProcessReceipt(prr)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error Processing Receipt")
		return
	}

	handleResponse(w, http.StatusOK, &responses.ProcessReceipt{ID: id})
}

func (rc *ReceiptsController) Points(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("Invalid request url reached Points Handler %v", r.URL))
		return
	}

	receipt, err := rc.store.GetReceiptById(id)
	if err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("error while looking up receipt: %v", err))
		return
	}

	if receipt == nil {
		handleError(w, http.StatusNotFound, "No receipt found for that id")
		return
	}

	handleResponse(w, http.StatusOK, &responses.Points{Points: receipt.Points})
}
